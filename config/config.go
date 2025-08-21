package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ahobsonsayers/twigots"
)

type Config struct {
	APIKey             string                    `json:"apiKey"`
	Country            twigots.Country           `json:"country"`
	RefetchTime        time.Duration             `json:"refetchTime"`
	Notification       NotificationConfig        `json:"notification"`
	GlobalTicketConfig GlobalTicketListingConfig `json:"global"`
	TicketConfigs      []TicketListingConfig     `json:"tickets"`
}

func (c Config) Validate() error {
	if c.APIKey == "" {
		return errors.New("api key must be set")
	}

	if c.Country.Value == "" {
		return errors.New("country must be set")
	}
	if !twigots.Countries.Contains(c.Country) {
		return fmt.Errorf("country '%s' is not valid", c.Country)
	}

	if c.RefetchTime <= 0 {
		c.RefetchTime = 1 * time.Minute // Default value
	}

	err := c.Notification.Validate()
	if err != nil {
		return fmt.Errorf("notification config is not valid: %w", err)
	}

	return nil
}

func (c Config) CombinedTicketListingConfigs() []TicketListingConfig {
	return CombineGlobalAndTicketListingConfigs(c.GlobalTicketConfig, c.TicketConfigs...)
}
