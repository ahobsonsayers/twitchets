package config

import (
	"errors"
	"fmt"

	"github.com/ahobsonsayers/twitchets/twickets"
)

type Config struct {
	APIKey        string             `json:"apiKey"`
	Country       twickets.Country   `json:"country"`
	Notification  NotificationConfig `json:"notification"`
	GlobalConfig  GlobalEventConfig  `json:"global"`
	TicketsConfig []TicketConfig     `json:"tickets"`
}

func (c Config) Validate() error {
	if c.APIKey == "" {
		return errors.New("api key must be set")
	}

	if c.Country.Value == "" {
		return errors.New("country must be set")
	}
	if !twickets.Countries.Contains(c.Country) {
		return fmt.Errorf("country '%s' is not valid", c.Country)
	}

	err := c.GlobalConfig.Validate()
	if err != nil {
		return fmt.Errorf("global config is not valid: %w", err)
	}

	for idx, ticketConfig := range c.TicketsConfig {
		err := ticketConfig.Validate()
		if err != nil {
			return fmt.Errorf("event config at index [%d] is no valid: %w", idx, err)
		}
	}

	return nil
}

func (c Config) CombineGlobalAndTicketConfig() []TicketConfig { // nolint: revive // TODO Remove nolint
	combinedConfigs := make([]TicketConfig, 0, len(c.TicketsConfig))
	for _, ticketConfig := range c.TicketsConfig {

		var combinedConfig TicketConfig
		combinedConfig.Event = ticketConfig.Event

		// Set name similarity
		if ticketConfig.EventSimilarity == nil {
			combinedConfig.EventSimilarity = &c.GlobalConfig.EventSimilarity
		} else if *ticketConfig.EventSimilarity > 0 {
			combinedConfig.EventSimilarity = ticketConfig.EventSimilarity
		}

		// Set regions
		if ticketConfig.Regions == nil {
			combinedConfig.Regions = c.GlobalConfig.Regions
		} else {
			combinedConfig.Regions = ticketConfig.Regions
		}

		// Set num tickets
		if ticketConfig.NumTickets == nil {
			combinedConfig.NumTickets = &c.GlobalConfig.NumTickets
		} else if *ticketConfig.NumTickets > 0 {
			combinedConfig.NumTickets = ticketConfig.NumTickets
		}

		// Set discount
		if ticketConfig.Discount == nil {
			combinedConfig.Discount = &c.GlobalConfig.Discount
		} else if *ticketConfig.Discount > 0 {
			combinedConfig.Discount = ticketConfig.Discount
		}

		// Set notification methods
		if ticketConfig.Notification == nil {
			combinedConfig.Notification = c.GlobalConfig.Notification
			if len(combinedConfig.Notification) == 0 {
				combinedConfig.Notification = NotificationTypes.Members()
			}
		} else {
			combinedConfig.Notification = ticketConfig.Notification
		}

		combinedConfigs = append(combinedConfigs, combinedConfig)
	}

	return combinedConfigs
}
