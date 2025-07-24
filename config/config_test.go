package config_test

import (
	"testing"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) { // nolint: revive
	configPath := test.ProjectDirectoryJoin(t, "test", "data", "config", "config.yaml")
	actualConfig, err := config.Load(configPath)
	require.NoError(t, err)

	country := twigots.CountryUnitedKingdom

	globalEventSimilarity := 0.75
	globalRegions := []twigots.Region{twigots.RegionLondon, twigots.RegionNorthWest}
	globalNumTickets := 2
	globalMaxTicketPrice := 25.0
	globalDiscount := 25.0

	expectedConfig := config.Config{
		APIKey:  "test",
		Country: country,
		GlobalTicketConfig: config.GlobalTicketListingConfig{
			EventSimilarity:       globalEventSimilarity,
			Regions:               globalRegions,
			MaxTicketPriceInclFee: globalMaxTicketPrice,
			NumTickets:            globalNumTickets,
			Min:                   globalDiscount,
		},
		Notification: config.NotificationConfig{
			Ntfy: &notification.NtfyConfig{
				Url:      "http://example.com",
				Topic:    "test",
				Username: "test",
				Password: "test",
			},
			Gotify: &notification.GotifyConfig{
				Url:   "http://example.com",
				Token: "test",
			},
			Telegram: &notification.TelegramConfig{
				Token:  "test",
				ChatId: 1234,
			},
		},
		TicketConfigs: []config.TicketListingConfig{
			{
				// Ticket with only event set
				Event: "Event 1",
			},
			{
				// Ticket with name similarity set
				Event:           "Event 2",
				EventSimilarity: lo.ToPtr(0.9),
			},
			{
				// Ticket with regions set
				Event:   "Event 3",
				Regions: []twigots.Region{twigots.RegionSouthWest},
			},
			{
				// Ticket with num tickets set
				Event:      "Event 4",
				NumTickets: lo.ToPtr(1),
			},
			{
				// Ticket with discount set
				Event:                 "Event 5",
				MaxTicketPriceInclFee: lo.ToPtr(15.0),
			},
			{
				// Ticket with discount set
				Event:       "Event 6",
				MinDiscount: lo.ToPtr(15.0),
			},
			{
				// Ticket with notification set
				Event:        "Event 7",
				Notification: []config.NotificationType{config.NotificationTypeNtfy},
			},
			{
				// Ticket with globals unset
				Event:                 "Event 8",
				EventSimilarity:       lo.ToPtr(-1.0),
				Regions:               []twigots.Region{},
				NumTickets:            lo.ToPtr(-1),
				MaxTicketPriceInclFee: lo.ToPtr(-1.0),
				MinDiscount:           lo.ToPtr(-1.0),
				Notification:          []config.NotificationType{},
			},
		},
	}

	require.Equal(t, expectedConfig, actualConfig)
}

func TestCombineConfigs(t *testing.T) { // nolint: revive
	configPath := test.ProjectDirectoryJoin(t, "test", "data", "config", "config.yaml")
	conf, err := config.Load(configPath)
	require.NoError(t, err)

	actualCombinedConfigs := conf.CombinedTicketListingConfigs()

	globalEventSimilarity := 0.75
	globalRegions := []twigots.Region{twigots.RegionLondon, twigots.RegionNorthWest}
	globalNumTickets := 2
	globalMaxTicketPrice := 25.0
	globalDiscount := 25.0

	expectedCombinedConfigs := []config.TicketListingConfig{
		{
			// Ticket with only event name set
			Event:                 "Event 1",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               globalRegions,
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           &globalDiscount,
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with event similarity set
			Event:                 "Event 2",
			EventSimilarity:       lo.ToPtr(0.90),
			Regions:               globalRegions,
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           &globalDiscount,
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with regions set
			Event:                 "Event 3",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               []twigots.Region{twigots.RegionSouthWest},
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           &globalDiscount,
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with num tickets set
			Event:                 "Event 4",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               globalRegions,
			NumTickets:            lo.ToPtr(1),
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           &globalDiscount,
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with max ticket price set
			Event:                 "Event 5",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               globalRegions,
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: lo.ToPtr(15.0),
			MinDiscount:           &globalDiscount,
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with discount set
			Event:                 "Event 6",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               globalRegions,
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           lo.ToPtr(15.0),
			Notification:          config.NotificationTypes.Members(),
		},
		{
			// Ticket with notification set
			Event:                 "Event 7",
			EventSimilarity:       &globalEventSimilarity,
			Regions:               globalRegions,
			NumTickets:            &globalNumTickets,
			MaxTicketPriceInclFee: &globalMaxTicketPrice,
			MinDiscount:           &globalDiscount,
			Notification:          []config.NotificationType{config.NotificationTypeNtfy},
		},
		{
			// Ticket with globals unset
			Event:                 "Event 8",
			EventSimilarity:       lo.ToPtr(-1.0),
			Regions:               []twigots.Region{},
			NumTickets:            lo.ToPtr(-1),
			MaxTicketPriceInclFee: lo.ToPtr(-1.0),
			MinDiscount:           lo.ToPtr(-1.0),
			Notification:          []config.NotificationType{},
		},
	}

	require.Equal(t, expectedCombinedConfigs, actualCombinedConfigs)
}
