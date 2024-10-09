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

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/config"
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

	// Twickets client
	twicketsClient := twickets.NewClient(nil)

	configPath := filepath.Join(cwd, "config.yaml")
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("config error:, %v", err)
	}

	// Notification Clients
	notificationClients, err := conf.Notification.Clients()
	if err != nil {
		log.Fatal(err)
	}

	// Event names
	eventNames := make([]string, 0, len(conf.TicketsConfig))
	for _, event := range conf.TicketsConfig {
		eventNames = append(eventNames, event.Event)
	}
	slog.Info(
		fmt.Sprintf("Monitoring: %s", strings.Join(eventNames, ", ")),
	)

	// Initial execution
	fetchAndProcessTickets(twicketsClient, conf, notificationClients)

	// Create ticker
	ticker := time.NewTicker(refetchTime)
	defer ticker.Stop()

	// Loop until exit
	exitChan := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets(twicketsClient, conf, notificationClients)
		case <-exitChan:
			return
		}
	}
}

func fetchAndProcessTickets(
	twicketsClient *twickets.Client,
	conf config.Config,
	notificationClients map[config.NotificationType]notification.Client,
) {
	checkTime := time.Now()
	defer func() {
		lastCheckTime = checkTime
	}()

	tickets, err := twicketsClient.FetchTickets(
		context.Background(),
		twickets.FetchTicketsInput{
			// Required
			APIKey:  conf.APIKey,
			Country: conf.Country,
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

	ticketConfigs := conf.CombineGlobalAndTicketConfig()
	for _, ticketConfig := range ticketConfigs {
		filter := ticketConfig.Filter()
		filteredTickets := tickets.Filter(filter)
		for _, ticket := range filteredTickets {
			slog.Info(
				"Found tickets for monitored event",
				"eventName", ticket.Event.Name,
				"numTickets", ticket.TicketQuantity,
				"ticketPrice", ticket.TotalTicketPrice().String(),
				"originalTicketPrice", ticket.OriginalTicketPrice().String(),
				"link", ticket.Link(),
			)

			for _, notificationType := range ticketConfig.Notification {

				notificationClient, ok := notificationClients[notificationType]
				if !ok {
					continue
				}

				err := notificationClient.SendTicketNotification(ticket)
				if err != nil {
					slog.Error(
						"Failed to send notification",
						"err", err,
					)
				}
			}

		}
	}
}
