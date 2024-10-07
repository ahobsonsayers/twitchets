package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/joho/godotenv"
)

const (
	maxNumTickets = 250
	refetchTime   = 1 * time.Minute
)

var lastCheckTime = time.Now()

func init() {
	_ = godotenv.Load()
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory:, %v", err)
	}

	configPath := filepath.Join(cwd, "config.yaml")
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("config error:, %v", err)
	}

	// Twickets client
	twicketsClient := twickets.NewClient(nil)

	// Notification Client
	notificationClient, err := notification.NewNtfyClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Event names
	eventNames := make([]string, 0, len(config.TicketsConfig))
	for _, event := range config.TicketsConfig {
		eventNames = append(eventNames, event.Event.Name)
	}
	slog.Info(
		fmt.Sprintf("Monitoring: %s", strings.Join(eventNames, ", ")),
	)

	// Initial execution
	fetchAndProcessTickets(config, twicketsClient, notificationClient)

	// Create ticker
	ticker := time.NewTicker(refetchTime)
	defer ticker.Stop()

	// Loop until exit
	exitChan := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets(config, twicketsClient, notificationClient)
		case <-exitChan:
			return
		}
	}
}

func fetchAndProcessTickets(
	config Config,
	twicketsClient *twickets.Client,
	notificationClient notification.Client,
) {
	checkTime := time.Now()
	defer func() {
		lastCheckTime = checkTime
	}()

	tickets, err := twicketsClient.FetchTickets(
		context.Background(),
		twickets.FetchTicketsInput{
			// Required
			APIKey:  config.APIKey,
			Country: config.Country,
			// Optional
			CreatedBefore: time.Now(),
			CreatedAfter:  lastCheckTime,
			NumTickets:    maxNumTickets,
		},
	)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if len(tickets) == maxNumTickets {
		slog.Warn("Fetched the max number of tickets allowed. It is possible tickets have been missed.")
	}

	filteredTickets := tickets.Filter(config.Filters())
	for _, ticket := range filteredTickets {
		slog.Info(
			"Found tickets for monitored event",
			"eventName", ticket.Event.Name,
			"numTickets", ticket.TicketQuantity,
			"ticketPrice", ticket.TotalTicketPrice().String(),
			"originalTicketPrice", ticket.OriginalTicketPrice().String(),
			"link", ticket.Link(),
		)

		err := notificationClient.SendTicketNotification(ticket)
		if err != nil {
			slog.Error(
				"Failed to send notification",
				"err", err,
			)
		}
	}
}
