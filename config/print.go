package config

import (
	"fmt"
	"strings"
)

func PrintTicketListingConfigs(configs []TicketListingConfig) {
	fmt.Println("Config:")
	fmt.Println()
	for _, config := range configs {
		PrintTicketListingConfig(config)
		fmt.Println()
	}
}

func PrintTicketListingConfig(config TicketListingConfig) { //nolint
	fmt.Printf("Event: %s\n", config.Event)

	if config.EventSimilarity == nil || *config.EventSimilarity <= 0.0 {
		fmt.Println("Event Similarity: Default (0.9)")
	} else {
		fmt.Printf("Event Similarity: %.2f%%\n", *config.EventSimilarity*100)
	}

	if len(config.Regions) == 0 {
		fmt.Println("Regions: Any")
	} else {

		// Get regions as a string
		regionStrings := make([]string, 0, len(config.Regions))
		for _, region := range config.Regions {
			regionStrings = append(regionStrings, region.Value)
		}
		regionsString := strings.Join(regionStrings, ", ")

		fmt.Printf("Regions: %s\n", regionsString)
	}

	if config.NumTickets == nil || *config.NumTickets <= 0 {
		fmt.Println("Number of Tickets: Any")
	} else {
		fmt.Printf("Number of Tickets: %d\n", *config.NumTickets)
	}

	if config.MinDiscount == nil || *config.MinDiscount <= 0.0 {
		fmt.Println("Discount: Any")
	} else {
		fmt.Printf("Discount: %.0f%%\n", *config.MinDiscount)
	}

	if len(config.Notification) == 0 {
		fmt.Println("Notification Types: All")
	} else {

		// Get notifications as a string
		notificationStrings := make([]string, 0, len(config.Notification))
		for _, notification := range config.Notification {
			notificationStrings = append(notificationStrings, notification.Value)
		}
		notificationsString := strings.Join(notificationStrings, ", ")

		fmt.Printf("Notifications: %s\n", notificationsString)
	}
}
