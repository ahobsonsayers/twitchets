package config

import (
	"errors"
	"fmt"

	"github.com/ahobsonsayers/twitchets/twickets"
)

type TicketConfig struct {
	Event           string             `json:"event"`
	EventSimilarity *float64           `json:"eventSimilarity"`
	Regions         []twickets.Region  `json:"regions"`
	NumTickets      *int               `json:"numTickets"`
	Discount        *float64           `json:"discount"`
	Notification    []NotificationType `json:"notification"`
}

func (t TicketConfig) Validate() error {
	if t.Event == "" {
		return errors.New("event name must be set")
	}

	for _, region := range t.Regions {
		if !twickets.Regions.Contains(region) {
			return fmt.Errorf("region '%s' is not valid", region)
		}
	}

	for _, notification := range t.Notification {
		if !NotificationTypes.Contains(notification) {
			return fmt.Errorf("notification type '%s' is not valid", notification)
		}
	}

	return nil
}

func (t TicketConfig) Filter() twickets.Filter {
	var filter twickets.Filter
	filter.Event = t.Event

	if t.EventSimilarity != nil {
		filter.EventSimilarity = *t.EventSimilarity
	}

	if t.Regions != nil {
		filter.Regions = t.Regions
	}

	if t.NumTickets != nil {
		filter.NumTickets = *t.NumTickets
	}

	if t.Discount != nil {
		filter.Discount = *t.Discount
	}

	return filter
}
