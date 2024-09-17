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

func getTestTickets(t *testing.T) twickets.Tickets {
	projectDirectory := testutils.GetProjectDirectory(t)
	feedJsonFilePath := filepath.Join(projectDirectory, "test", "assets", "feed.json")

	feedJsonFile, err := os.Open(feedJsonFilePath)
	require.NoError(t, err)
	feedJson, err := io.ReadAll(feedJsonFile)
	require.NoError(t, err)

	tickets, err := twickets.UnmarshalTwicketsFeedJson(feedJson)
	require.NoError(t, err)

	return tickets
}

func TestUnmarshalFeedJson(t *testing.T) {
	tickets := getTestTickets(t)

	require.Len(t, tickets, 4)

	require.Equal(t, "Foo Fighters", tickets[0].Event.Name)
	require.Len(t, tickets[0].Event.Lineup, 3)
	require.Equal(t, "Foo Fighters", tickets[0].Event.Lineup[0].Artist.Name)
	require.Equal(t, "Wet Leg", tickets[0].Event.Lineup[1].Artist.Name)
	require.Equal(t, "Shame", tickets[0].Event.Lineup[2].Artist.Name)
	require.Equal(t, "London Stadium", tickets[0].Event.Venue.Name)
	require.Equal(t, "Foo Fighters - Everything Or Nothing At All Tour", tickets[0].Tour.Name)
	require.Equal(t, 3, tickets[0].TicketQuantity)
	require.Equal(t, "£180.00", tickets[0].TicketsPrice.String())
	require.Equal(t, "£38.25", tickets[0].TwicketsFee.String())
	require.Equal(t, "£255.00", tickets[0].OriginalTotalPrice.String())

	require.Equal(t, "Mean Girls", tickets[1].Event.Name)
	require.Empty(t, tickets[1].Event.Lineup)
	require.Equal(t, "Savoy Theatre", tickets[1].Event.Venue.Name)
	require.Equal(t, "Mean Girls", tickets[1].Tour.Name)
	require.Equal(t, 2, tickets[1].TicketQuantity)
	require.Equal(t, "£130.00", tickets[1].TicketsPrice.String())
	require.Equal(t, "£18.20", tickets[1].TwicketsFee.String())
	require.Equal(t, "£130.00", tickets[1].OriginalTotalPrice.String())

	require.Equal(t, "South Africa v Wales", tickets[2].Event.Name)
	require.Empty(t, tickets[2].Event.Lineup)
	require.Equal(t, "Twickenham Stadium", tickets[2].Event.Venue.Name)
	require.Equal(t, "South Africa v Wales", tickets[2].Tour.Name)
	require.Equal(t, 4, tickets[2].TicketQuantity)
	require.Equal(t, "£380.00", tickets[2].TicketsPrice.String())
	require.Equal(t, "£53.20", tickets[2].TwicketsFee.String())
	require.Equal(t, "£380.00", tickets[2].OriginalTotalPrice.String())

	require.Equal(t, "Download Festival 2024", tickets[3].Event.Name)
	require.Empty(t, tickets[3].Event.Lineup)
	require.Equal(t, "Donington Park", tickets[3].Event.Venue.Name)
	require.Equal(t, "Download Festival 2024", tickets[3].Tour.Name)
	require.Equal(t, 1, tickets[3].TicketQuantity)
	require.Equal(t, "£280.00", tickets[3].TicketsPrice.String())
	require.Equal(t, "£30.80", tickets[3].TwicketsFee.String())
	require.Equal(t, "£322.00", tickets[3].OriginalTotalPrice.String())
}

func TestGetEventById(t *testing.T) {
	tickets := getTestTickets(t)
	ticket := tickets.GetById("156783487261837")
	require.NotNil(t, ticket)
}

func TestDiscount(t *testing.T) {
	tickets := getTestTickets(t)
	discountString := tickets[0].DiscountString()
	require.Equal(t, "14.41%", discountString)
}

// TestFilterName tests we find tickets for event names that are
// not specified quite right
func TestFilterName(t *testing.T) {
	strangerThingsActual := "Stranger Things: The First Shadow"
	strangerThingsTest := "Stranger Things"

	backToTheFutureActual := "Back To The Future: The Musical"
	backToTheFutureTest := "Back To The Future"

	harryPotterActual := "Harry Potter & The Cursed Child Parts 1 & 2"
	harryPotterTest := "Harry Potter and the Cursed Child"

	tickets := twickets.Tickets{
		{Event: twickets.Event{Name: strangerThingsActual}},
		{Event: twickets.Event{Name: backToTheFutureActual}},
		{Event: twickets.Event{Name: harryPotterActual}},
	}

	// Stranger Things
	gotTickets := tickets.Filter(twickets.TicketFilter{
		EventNames: []string{strangerThingsTest},
	})
	require.Len(t, gotTickets, 1)
	require.Equal(t, strangerThingsActual, gotTickets[0].Event.Name)

	// Back to the Future
	gotTickets = tickets.Filter(twickets.TicketFilter{
		EventNames: []string{backToTheFutureTest},
	})
	require.Len(t, gotTickets, 1)
	require.Equal(t, backToTheFutureActual, gotTickets[0].Event.Name)

	// Harry Potter
	gotTickets = tickets.Filter(twickets.TicketFilter{
		EventNames: []string{harryPotterTest},
	})
	require.Len(t, gotTickets, 1)
	require.Equal(t, harryPotterActual, gotTickets[0].Event.Name)
}
