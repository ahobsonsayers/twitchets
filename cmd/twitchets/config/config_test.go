package config_test

import (
	"testing"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/config"
	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) { // nolint: revive
	configPath := testutils.ProjectDirectoryJoin(t, "test", "testdata", "config", "config.yaml")
	actualConfig, err := config.Load(configPath)
	require.NoError(t, err)

	country := twickets.CountryUnitedKingdom

	globalEventSimilarity := 75.0
	globalRegions := []twickets.Region{twickets.RegionLondon, twickets.RegionNorthWest}
	globalNumTickets := 2
	globalDiscount := 25.0

	expectedConfig := config.Config{
		APIKey:  "test",
		Country: country,
		GlobalConfig: config.GlobalEventConfig{
			EventSimilarity: globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      globalNumTickets,
			Discount:        globalDiscount,
		},
		Notification: config.NotificationConfig{
			Ntfy: &notification.NtfyConfig{
				Url:      "example.com",
				Topic:    "test",
				Username: "test",
				Password: "test",
			},
			Telegram: &notification.TelegramConfig{
				APIToken: "test",
				ChatId:   1234,
			},
		},
		TicketsConfig: []config.TicketConfig{
			{
				// Ticket with only event set
				Event: "Event 1",
			},
			{
				// Ticket with name similarity set
				Event:           "Event 2",
				EventSimilarity: lo.ToPtr(90.0),
			},
			{
				// Ticket with regions set
				Event:   "Event 3",
				Regions: []twickets.Region{twickets.RegionSouthWest},
			},
			{
				// Ticket with num tickets set
				Event:      "Event 4",
				NumTickets: lo.ToPtr(1),
			},
			{
				// Ticket with discount set
				Event:    "Event 5",
				Discount: lo.ToPtr(15.0),
			},
			{
				// Ticket with notification set
				Event:        "Event 6",
				Notification: []config.NotificationType{config.NotificationTypeNtfy},
			},
			{
				// Ticket with globals unset
				Event:           "Event 7",
				EventSimilarity: lo.ToPtr(-1.0),
				Regions:         []twickets.Region{},
				NumTickets:      lo.ToPtr(-1),
				Discount:        lo.ToPtr(-1.0),
				Notification:    []config.NotificationType{},
			},
		},
	}

	require.EqualValues(t, expectedConfig, actualConfig)
}

func TestCombineConfigs(t *testing.T) { // nolint: revive
	configPath := testutils.ProjectDirectoryJoin(t, "test", "testdata", "config", "config.yaml")
	conf, err := config.Load(configPath)
	require.NoError(t, err)

	actualCombinedConfigs := conf.CombineGlobalAndTicketConfig()

	globalEventSimilarity := 75.0
	globalRegions := []twickets.Region{twickets.RegionLondon, twickets.RegionNorthWest}
	globalNumTickets := 2
	globalDiscount := 25.0

	expectedCombinedConfigs := []config.TicketConfig{
		{
			// Ticket with only event name set
			Event:           "Event 1",
			EventSimilarity: &globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      &globalNumTickets,
			Discount:        &globalDiscount,
			Notification:    config.NotificationTypes.Members(),
		},
		{
			// Ticket with event similarity set
			Event:           "Event 2",
			EventSimilarity: lo.ToPtr(90.0),
			Regions:         globalRegions,
			NumTickets:      &globalNumTickets,
			Discount:        &globalDiscount,
			Notification:    config.NotificationTypes.Members(),
		},
		{
			// Ticket with regions set
			Event:           "Event 3",
			EventSimilarity: &globalEventSimilarity,
			Regions:         []twickets.Region{twickets.RegionSouthWest},
			NumTickets:      &globalNumTickets,
			Discount:        &globalDiscount,
			Notification:    config.NotificationTypes.Members(),
		},
		{
			// Ticket with num tickets set
			Event:           "Event 4",
			EventSimilarity: &globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      lo.ToPtr(1),
			Discount:        &globalDiscount,
			Notification:    config.NotificationTypes.Members(),
		},
		{
			// Ticket with discount set
			Event:           "Event 5",
			EventSimilarity: &globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      &globalNumTickets,
			Discount:        lo.ToPtr(15.0),
			Notification:    config.NotificationTypes.Members(),
		},
		{
			// Ticket with notification set
			Event:           "Event 6",
			EventSimilarity: &globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      &globalNumTickets,
			Discount:        &globalDiscount,
			Notification:    []config.NotificationType{config.NotificationTypeNtfy},
		},
		{
			// Ticket with globals unset
			Event:           "Event 7",
			EventSimilarity: nil,
			Regions:         []twickets.Region{},
			NumTickets:      nil,
			Discount:        nil,
			Notification:    []config.NotificationType{},
		},
	}

	require.EqualValues(t, expectedCombinedConfigs, actualCombinedConfigs)
}
