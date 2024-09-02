package twickets_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalFeedJson(t *testing.T) {
	projectDirectory := testutils.GetProjectDirectory(t)
	feedJsonFilePath := filepath.Join(projectDirectory, "test", "assets", "feed.json")

	feedJsonFile, err := os.Open(feedJsonFilePath)
	require.NoError(t, err)
	feedJson, err := io.ReadAll(feedJsonFile)
	require.NoError(t, err)

	tickets, err := twickets.UnmarshalTwicketsFeedJson(feedJson)
	require.NoError(t, err)

	require.Len(t, tickets, 4)

	require.Equal(t, "Foo Fighters", tickets[0].Event.Name)
	require.Len(t, tickets[0].Event.Lineup, 3)
	require.Equal(t, "Foo Fighters", tickets[0].Event.Lineup[0].Artist.Name)
	require.Equal(t, "Wet Leg", tickets[0].Event.Lineup[1].Artist.Name)
	require.Equal(t, "Shame", tickets[0].Event.Lineup[2].Artist.Name)
	require.Equal(t, "London Stadium", tickets[0].Event.Venue.Name)
	require.Equal(t, "Foo Fighters - Everything Or Nothing At All Tour", tickets[0].Tour.Name)
	require.Equal(t, 3, tickets[0].TicketQuantity)
	require.Equal(t, "£255.00", tickets[0].TotalSellingPrice.String())
	require.Equal(t, "£38.25", tickets[0].TotalTwicketsFee.String())
	require.Equal(t, "£255.00", tickets[0].FaceValuePrice.String())

	require.Equal(t, "Mean Girls", tickets[1].Event.Name)
	require.Empty(t, tickets[1].Event.Lineup)
	require.Equal(t, "Savoy Theatre", tickets[1].Event.Venue.Name)
	require.Equal(t, "Mean Girls", tickets[1].Tour.Name)
	require.Equal(t, 2, tickets[1].TicketQuantity)
	require.Equal(t, "£130.00", tickets[1].TotalSellingPrice.String())
	require.Equal(t, "£18.20", tickets[1].TotalTwicketsFee.String())
	require.Equal(t, "£130.00", tickets[1].FaceValuePrice.String())

	require.Equal(t, "South Africa v Wales", tickets[2].Event.Name)
	require.Empty(t, tickets[2].Event.Lineup)
	require.Equal(t, "Twickenham Stadium", tickets[2].Event.Venue.Name)
	require.Equal(t, "South Africa v Wales", tickets[2].Tour.Name)
	require.Equal(t, 4, tickets[2].TicketQuantity)
	require.Equal(t, "£380.00", tickets[2].TotalSellingPrice.String())
	require.Equal(t, "£53.20", tickets[2].TotalTwicketsFee.String())
	require.Equal(t, "£380.00", tickets[2].FaceValuePrice.String())

	require.Equal(t, "Download Festival 2024", tickets[3].Event.Name)
	require.Empty(t, tickets[3].Event.Lineup)
	require.Equal(t, "Donington Park", tickets[3].Event.Venue.Name)
	require.Equal(t, "Download Festival 2024", tickets[3].Tour.Name)
	require.Equal(t, 1, tickets[3].TicketQuantity)
	require.Equal(t, "£280.00", tickets[3].TotalSellingPrice.String())
	require.Equal(t, "£30.80", tickets[3].TotalTwicketsFee.String())
	require.Equal(t, "£322.00", tickets[3].FaceValuePrice.String())
}
