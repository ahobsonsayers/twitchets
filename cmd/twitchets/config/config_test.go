package config_test

import (
	"testing"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/config"
	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	configPath := testutils.ProjectDirectoryJoin(t, "test", "assets", "config", "config.yaml")
	actualConfig, err := config.LoadConfig(configPath)
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
		TicketsConfig: []config.TicketConfig{
			{
				// Ticket with only event set
				Event: config.Event{Name: "Event 1"},
			},
			{
				// Ticket with name similarity set
				Event:           config.Event{Name: "Event 2"},
				EventSimilarity: lo.ToPtr(90.0),
			},
			{
				// Ticket with regions set
				Event:   config.Event{Name: "Event 3"},
				Regions: []twickets.Region{twickets.RegionSouthWest},
			},
			{
				// Ticket with num tickets set
				Event:      config.Event{Name: "Event 4"},
				NumTickets: lo.ToPtr(1),
			},
			{
				// Ticket with discount set
				Event:    config.Event{Name: "Event 5"},
				Discount: lo.ToPtr(15.0),
			},
			{
				// Ticket with globals unset
				Event:           config.Event{Name: "Event 6"},
				EventSimilarity: lo.ToPtr(-1.0),
				Regions:         []twickets.Region{},
				NumTickets:      lo.ToPtr(-1),
				Discount:        lo.ToPtr(-1.0),
			},
		},
	}

	require.EqualValues(t, expectedConfig, actualConfig)
}

func TestConfigFilters(t *testing.T) {
	configPath := testutils.ProjectDirectoryJoin(t, "test", "assets", "config", "config.yaml")
	config, err := config.LoadConfig(configPath)
	require.NoError(t, err)

	actualFilters := config.Filters()

	globalEventSimilarity := 75.0
	globalRegions := []twickets.Region{twickets.RegionLondon, twickets.RegionNorthWest}
	globalNumTickets := 2
	globalDiscount := 25.0

	expectedFilters := []twickets.Filter{
		{
			// Ticket with only event name set
			Event:           "Event 1",
			EventSimilarity: globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      globalNumTickets,
			Discount:        globalDiscount,
		},
		{
			// Ticket with event similarity set
			Event:           "Event 2",
			EventSimilarity: 90.0,
			Regions:         globalRegions,
			NumTickets:      globalNumTickets,
			Discount:        globalDiscount,
		},
		{
			// Ticket with regions set
			Event:           "Event 3",
			EventSimilarity: globalEventSimilarity,
			Regions:         []twickets.Region{twickets.RegionSouthWest},
			NumTickets:      globalNumTickets,
			Discount:        globalDiscount,
		},
		{
			// Ticket with num tickets set
			Event:           "Event 4",
			EventSimilarity: globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      1,
			Discount:        globalDiscount,
		},
		{
			// Ticket with discount set
			Event:           "Event 5",
			EventSimilarity: globalEventSimilarity,
			Regions:         globalRegions,
			NumTickets:      globalNumTickets,
			Discount:        15.0,
		},
		{
			// Ticket with globals unset
			Event:           "Event 6",
			EventSimilarity: 0.0,
			Regions:         []twickets.Region{},
			NumTickets:      0,
			Discount:        0.0,
		},
	}

	require.EqualValues(t, expectedFilters, actualFilters)
}
