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

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
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
	client := twigots.NewClient()

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
	fetchAndProcessTickets(client, conf, notificationClients)

	// Create ticker
	ticker := time.NewTicker(refetchTime)
	defer ticker.Stop()

	// Loop until exit
	exitChan := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets(client, conf, notificationClients)
		case <-exitChan:
			return
		}
	}
}

func fetchAndProcessTickets(
	client *twigots.Client,
	conf config.Config,
	notificationClients map[config.NotificationType]notification.Client,
) {
	checkTime := time.Now()
	defer func() {
		lastCheckTime = checkTime
	}()

	listings, err := client.FetchTicketListings(
		context.Background(),
		twigots.FetchTicketListingsInput{
			// Required
			APIKey:  conf.APIKey,
			Country: twigots.CountryUnitedKingdom,
			// Optional
			CreatedBefore: time.Now(),
			CreatedAfter:  lastCheckTime,
			MaxNumber:     maxNumTickets,
		},
	)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if len(listings) == maxNumTickets {
		slog.Warn("Fetched the max number of tickets allowed. It is possible tickets have been missed.")
	}

	ticketConfigs := conf.CombineGlobalAndTicketConfig()
	for _, ticketConfig := range ticketConfigs {
		filter := ticketConfig.Filter()
		filteredListings, err := listings.Filter(
			twigots.Filter{
				Event:           filter.Event,
				EventSimilarity: filter.EventSimilarity,
				Regions:         filter.Regions,
				NumTickets:      filter.NumTickets,
				MinDiscount:     filter.MinDiscount,
				CreatedAfter:    filter.CreatedAfter,
			},
		)
		if err != nil {
			slog.Error(
				"Failed to filter listings",
				"err", err,
			)
			continue
		}

		for _, listing := range filteredListings {
			slog.Info(
				"Found tickets for monitored event",
				"monitoredEventString", ticketConfig.Event,
				"matchedEventName", listing.Event.Name,
				"numTickets", listing.NumTickets,
				"ticketPrice", listing.TotalPriceInclFee().String(),
				"originalTicketPrice", listing.OriginalTicketPrice().String(),
				"link", listing.URL(),
			)

			for _, notificationType := range ticketConfig.Notification {
				notificationClient, ok := notificationClients[notificationType]
				if !ok {
					continue
				}

				err := notificationClient.SendTicketNotification(listing)
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
