package notification

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/config"
)

var (
	//go:embed templates/message.tmpl.md
	messageTemplateFS embed.FS
	messageTemplate   *template.Template
)

func init() {
	var err error
	messageTemplate, err = template.ParseFS(messageTemplateFS, "templates/message.tmpl.md")
	if err != nil {
		log.Fatalf("failed to read notification message template: %v", err)
	}
}

type Client interface {
	SendTicketNotification(twigots.TicketListing) error
}

type MessageTemplateData struct {
	// Header
	Event string

	// Body
	Date                string
	Time                string
	Venue               string
	Location            string
	TicketType          string // Standing, Stalls etc.
	NumTickets          int
	TotalTicketPrice    string
	TotalPrice          string
	OriginalTicketPrice string
	OriginalTotalPrice  string
	Discount            string
	AcceptsOffers       bool

	// Footer
	Link string
}

type renderMessageConfig struct {
	includeHeader bool
	includeFooter bool
}

func newRenderMessageConfig(options ...RenderMessageOption) renderMessageConfig {
	conf := new(renderMessageConfig)
	for _, option := range options {
		if option != nil {
			option(conf)
		}
	}
	return *conf
}

type RenderMessageOption func(*renderMessageConfig)

// Whether to include event name header in message
func WithHeader() RenderMessageOption {
	return func(o *renderMessageConfig) {
		o.includeHeader = true
	}
}

// Whether to include buy link footer in message
func WithFooter() RenderMessageOption {
	return func(o *renderMessageConfig) {
		o.includeFooter = true
	}
}

func RenderMessage(ticket twigots.TicketListing, options ...RenderMessageOption) (string, error) {
	config := newRenderMessageConfig(options...)

	templateData := MessageTemplateData{
		Date:                ticket.Event.Date.Format("Monday 2 January 2006"),
		Time:                ticket.Event.Time.Format("3:04pm"),
		Venue:               ticket.Event.Venue.Name,
		Location:            ticket.Event.Venue.Location.Name,
		TicketType:          ticket.TicketType,
		NumTickets:          ticket.NumTickets,
		TotalTicketPrice:    ticket.TicketPriceInclFee().String(),
		TotalPrice:          ticket.TotalPriceInclFee().String(),
		OriginalTicketPrice: ticket.OriginalTicketPrice().String(),
		OriginalTotalPrice:  ticket.OriginalTotalPrice.String(),
		Discount:            ticket.DiscountString(),
		AcceptsOffers:       ticket.SellerWillConsiderOffers,
	}

	// Add optional header and footers
	if config.includeHeader {
		templateData.Event = ticket.Event.Name
	}

	if config.includeFooter {
		templateData.Link = ticket.URL()
	}

	var buffer bytes.Buffer
	err := messageTemplate.Execute(&buffer, templateData)
	if err != nil {
		return "", fmt.Errorf("failed to render notification message template:, %w", err)
	}

	message := buffer.String()
	message = strings.TrimSpace(message)

	return message, nil
}

func GetNotificationClients(conf config.NotificationConfig) (map[config.NotificationType]Client, error) {
	clients := map[config.NotificationType]Client{}

	if conf.Ntfy != nil {
		ntfyClient, err := NewNtfyClient(*conf.Ntfy)
		if err != nil {
			return nil, fmt.Errorf("failed to setup ntfy client: %w", err)
		}

		clients[config.NotificationTypeNtfy] = ntfyClient
	}

	if conf.Gotify != nil {
		gotifyClient, err := NewGotifyClient(*conf.Gotify)
		if err != nil {
			return nil, fmt.Errorf("failed to setup gotify client: %w", err)
		}

		clients[config.NotificationTypeGotify] = gotifyClient
	}

	if conf.Telegram != nil {
		telegramClient, err := NewTelegramClient(*conf.Telegram)
		if err != nil {
			return nil, fmt.Errorf("failed to setup telegram client: %w", err)
		}

		clients[config.NotificationTypeTelegram] = telegramClient
	}

	return clients, nil
}
