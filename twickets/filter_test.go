package twickets_test

import (
	"testing"
	"time"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/stretchr/testify/require"
)

// TestFilterName tests we find tickets for event names that are
// not specified quite right
func TestFilterName(t *testing.T) {
	strangerThingsAsked := "Stranger Things"
	strangerThingsGot := "Stranger Things: The First Shadow"

	backToTheFutureAsked := "Back To The Future"
	backToTheFutureGot := "Back To The Future: The Musical"

	harryPotterAsked := "Harry Potter and the Cursed Child"
	harryPotterGot := "Harry Potter & The Cursed Child Parts 1 & 2"

	wizardOfOzAsked := "The Who"
	wizardOfOzGot := "The The" // This shouldn't match

	gotTickets := twickets.Tickets{
		{Event: twickets.Event{Name: strangerThingsGot}},
		{Event: twickets.Event{Name: backToTheFutureGot}},
		{Event: twickets.Event{Name: harryPotterGot}},
		{Event: twickets.Event{Name: wizardOfOzGot}},
	}

	// Stranger Things
	filteredTickets := gotTickets.Filter(twickets.Filter{
		Event: strangerThingsAsked,
	})
	require.Len(t, filteredTickets, 1)
	require.Equal(t, strangerThingsGot, filteredTickets[0].Event.Name)

	// Back to the Future
	filteredTickets = gotTickets.Filter(twickets.Filter{
		Event: backToTheFutureAsked,
	})
	require.Len(t, filteredTickets, 1)
	require.Equal(t, backToTheFutureGot, filteredTickets[0].Event.Name)

	// Harry Potter
	filteredTickets = gotTickets.Filter(twickets.Filter{
		Event: harryPotterAsked,
	})
	require.Len(t, filteredTickets, 1)
	require.Equal(t, harryPotterGot, filteredTickets[0].Event.Name)

	// Wizard of Oz
	filteredTickets = gotTickets.Filter(twickets.Filter{
		Event: wizardOfOzAsked,
	})
	require.Empty(t, filteredTickets)
}

func TestFilterTicketsToCreatedAfter(t *testing.T) {
	currentTime := time.Now()
	tickets := twickets.Tickets{
		{CreatedAt: twickets.UnixTime{currentTime.Add(-1 * time.Minute)}},
		{CreatedAt: twickets.UnixTime{currentTime.Add(-2 * time.Minute)}},
		{CreatedAt: twickets.UnixTime{currentTime.Add(-4 * time.Minute)}},
		{CreatedAt: twickets.UnixTime{currentTime.Add(-5 * time.Minute)}},
	}

	filteredTickets := twickets.FilterTicketsCreatedAfter(
		tickets,
		currentTime.Add(-3*time.Minute),
	)

	require.Equal(t, tickets[0:2], filteredTickets)
}
