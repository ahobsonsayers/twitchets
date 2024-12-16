package config

import (
	"github.com/ahobsonsayers/twigots"
)

// GlobalEventConfig is config that applies to all events,
// unless an event explicitly overwrites its.
// Country is required.
type GlobalEventConfig struct {
	EventSimilarity float64            `json:"eventSimilarity"`
	Regions         []twigots.Region   `json:"regions"`
	NumTickets      int                `json:"numTickets"`
	Discount        float64            `json:"discount"`
	Notification    []NotificationType `json:"notification"`
}

func (c GlobalEventConfig) Validate() error {
	// Reuse the filter validation logic
	filter := twigots.Filter{
		Event:           "global", // Event must be be set - this is arbitrary
		EventSimilarity: c.EventSimilarity,
		Regions:         c.Regions,
		NumTickets:      c.NumTickets,
		MinDiscount:     c.Discount,
	}
	err := filter.Validate()
	if err != nil {
		return err
	}

	return nil
}
