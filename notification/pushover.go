package notification

import (
	"fmt"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/gregdel/pushover"
)

type PushoverClient struct {
	app       *pushover.Pushover
	recipient *pushover.Recipient
	config    config.PushoverConfig
}

var _ Client = PushoverClient{}

func (p PushoverClient) SendTicketNotification(ticket twigots.TicketListing) error {
	notificationMessage, err := RenderMessage(ticket)
	if err != nil {
		return err
	}

	message := pushover.NewMessage(notificationMessage)
	message.Title = ticket.Event.Name
	message.URL = ticket.URL()
	message.URLTitle = "Buy Tickets"

	// Apply optional configuration
	if p.config.Device != nil && *p.config.Device != "" {
		message.DeviceName = *p.config.Device
	}

	if p.config.Sound != nil && *p.config.Sound != "" {
		message.Sound = *p.config.Sound
	}

	if p.config.Priority != nil {
		message.Priority = *p.config.Priority
	}

	_, err = p.app.SendMessage(message, p.recipient)
	if err != nil {
		return fmt.Errorf("failed to send pushover notification: %w", err)
	}

	return nil
}

func NewPushoverClient(conf config.PushoverConfig) (PushoverClient, error) {
	app := pushover.New(conf.Token)
	recipient := pushover.NewRecipient(conf.User)

	return PushoverClient{
		app:       app,
		recipient: recipient,
		config:    conf,
	}, nil
}
