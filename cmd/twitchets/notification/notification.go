package notification

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/ahobsonsayers/twitchets/twickets"
)

var (
	//go:embed template/message.tmpl.md
	messageTemplateFS embed.FS
	messageTemplate   *template.Template
)

func init() {
	var err error
	messageTemplate, err = template.ParseFS(messageTemplateFS, "template/message.tmpl.md")
	if err != nil {
		log.Fatalf("failed to read notification message template: %v", err)
	}
}

type Client interface {
	SendTicketNotification(twickets.Ticket) error
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
	Discount            float64

	// Footer
	Link string
}

type renderMessageConfig struct {
	includeHeader bool
	includeFooter bool
}

func (c *renderMessageConfig) applyOptions(options ...RenderMessageOption) {
	for _, option := range options {
		option(c)
	}
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

func RenderMessage(ticket twickets.Ticket, options ...RenderMessageOption) (string, error) {
	config := &renderMessageConfig{}
	config.applyOptions(options...)

	templateData := MessageTemplateData{
		Date:                ticket.Event.Date.Format("Monday 2 January 2006"),
		Time:                ticket.Event.Time.Format("3:04pm"),
		Venue:               ticket.Event.Venue.Name,
		Location:            ticket.Event.Venue.Location.Name,
		TicketType:          ticket.TicketType,
		NumTickets:          ticket.TicketQuantity,
		TotalTicketPrice:    ticket.TotalTicketPrice().String(),
		TotalPrice:          ticket.TotalPrice().String(),
		OriginalTicketPrice: ticket.OriginalTicketPrice().String(),
		OriginalTotalPrice:  ticket.OriginalTotalPrice.String(),
		Discount:            ticket.Discount(),
	}

	// Add optional header and footers
	if config.includeHeader {
		templateData.Event = ticket.Event.Name
	}

	if config.includeFooter {
		templateData.Link = ticket.Link()
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
