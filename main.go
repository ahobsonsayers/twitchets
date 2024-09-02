package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/samber/lo"
)

var (
	lastCheckTime       = time.Time{}
	monitoredEventNames = []string{
		// Theatre
		"Back to the Future",
		"Frozen",
		"Hamilton",
		"Harry Potter & the Cursed Child",
		"Kiss Me Kate",
		"Lion King",
		"Matilda",
		"Mean Girls",
		"Moulin Rouge",
		"Starlight Express",
		"Stranger Things",
		"The Phantom Opera",
		"The Wizard of Oz",
		// Gigs
		"Coldplay",
		"Taylor Swift",
	}
)

func main() {
	slog.Info(
		fmt.Sprintf("Monitoring: %s", strings.Join(monitoredEventNames, ", ")),
	)

	// Initial execution
	fetchAndProcessTickets()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	exitChan := make(chan struct{})

	// Loop until exit
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets()
		case <-exitChan:
			return
		}
	}
}

func fetchAndProcessTickets() {
	checkTime := time.Now()

	tickets, err := twickets.FetchLatestTickets(
		context.Background(),
		twickets.GetTicketsInput{
			Country: twickets.CountryUnitedKingdom,
			// Regions:    []twickets.Region{twickets.RegionLondon},
			MaxNumber:  10,
			BeforeTime: checkTime,
		},
	)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	processTickets(tickets)
	lastCheckTime = checkTime
}

func processTickets(tickets []twickets.Ticket) {
	for _, ticket := range lo.Reverse(tickets) {
		if ticket.CreatedAt.Before(lastCheckTime) {
			continue
		}

		for _, eventName := range monitoredEventNames {

			isMonitored := fuzzy.MatchNormalizedFold(eventName, ticket.Event.Name) ||
				fuzzy.MatchNormalizedFold(ticket.Event.Name, eventName)

			if isMonitored {
				slog.Info(
					"found tickets for monitored event",
					"name", ticket.Event.Name,
					"tickets", ticket.TicketQuantity,
					"ticketCost", ticket.TotalSellingPrice.PerString(ticket.TicketQuantity),
					"totalCost", ticket.TotalSellingPrice.String(),
				)
			}
		}
	}
}
