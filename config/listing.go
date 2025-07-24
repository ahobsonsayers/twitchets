package config

import (
	"github.com/ahobsonsayers/twigots"
)

// GlobalTicketListingConfig represents configuration settings that apply to all ticket listings
// unless explicitly overridden by a specific ticket configuration.
type GlobalTicketListingConfig struct {
	EventSimilarity       float64            `json:"eventSimilarity"`
	Regions               []twigots.Region   `json:"regions"`
	NumTickets            int                `json:"numTickets"`
	Min                   float64            `json:"discount"`
	MaxTicketPriceInclFee float64            `json:"maxTicketPrice"`
	Notification          []NotificationType `json:"notification"`
}

// TicketListingConfig represents configuration for a specific ticket listing
type TicketListingConfig struct {
	Event                 string             `json:"event"`
	EventSimilarity       *float64           `json:"eventSimilarity"`
	Regions               []twigots.Region   `json:"regions"`
	NumTickets            *int               `json:"numTickets"`
	MinDiscount           *float64           `json:"discount"`
	MaxTicketPriceInclFee *float64           `json:"maxTicketPrice"`
	Notification          []NotificationType `json:"notification"`
}

// CombineGlobalAndTicketConfigs merges global and specific ticket listing configurations.
//
// It returns a slice of TicketListingConfig where each configuration has global ticket listing configuration
// merged with specific ticket listing configuration.
//
// If specific ticket listing configuration is provided, it takes precedence over the global configuration.
func CombineGlobalAndTicketListingConfigs( //nolint: revive
	globalConfig GlobalTicketListingConfig,
	configs ...TicketListingConfig,
) []TicketListingConfig {
	combinedConfigs := make([]TicketListingConfig, 0, len(configs))

	for _, config := range configs {

		var combinedConfig TicketListingConfig
		// Set event name
		combinedConfig.Event = config.Event

		// Set event similarity, using global if not specified
		if config.EventSimilarity == nil {
			combinedConfig.EventSimilarity = &globalConfig.EventSimilarity
		} else {
			combinedConfig.EventSimilarity = config.EventSimilarity
		}

		// Set regions, using global if not specified
		if config.Regions == nil {
			combinedConfig.Regions = globalConfig.Regions
		} else {
			combinedConfig.Regions = config.Regions
		}

		// Set number of tickets, using global if not specified
		if config.NumTickets == nil {
			combinedConfig.NumTickets = &globalConfig.NumTickets
		} else {
			combinedConfig.NumTickets = config.NumTickets
		}

		// Set discount, using global if not specified
		if config.MinDiscount == nil {
			combinedConfig.MinDiscount = &globalConfig.Min
		} else {
			combinedConfig.MinDiscount = config.MinDiscount
		}

		// Set max ticket price including fee, using global if not specified
		if config.MaxTicketPriceInclFee == nil {
			combinedConfig.MaxTicketPriceInclFee = &globalConfig.MaxTicketPriceInclFee
		} else {
			combinedConfig.MaxTicketPriceInclFee = config.MaxTicketPriceInclFee
		}

		// Set notifications, using global if not specified
		// Default to all notification types if both are empty
		if config.Notification == nil {
			combinedConfig.Notification = globalConfig.Notification
			if len(combinedConfig.Notification) == 0 {
				combinedConfig.Notification = NotificationTypes.Members()
			}
		} else {
			combinedConfig.Notification = config.Notification
		}

		combinedConfigs = append(combinedConfigs, combinedConfig)
	}

	return combinedConfigs
}
