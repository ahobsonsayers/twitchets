package main

import (
	"context"
	"log"
	"log/slog"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twigots/filter"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
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

	// Load config
	configPath := filepath.Join(cwd, "config.yaml")
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("config error:, %v", err)
	}

	// Create twickets client
	client, err := twigots.NewClient(conf.APIKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create notification clients
	notificationClients, err := conf.Notification.Clients()
	if err != nil {
		log.Fatal(err)
	}

	// Get combined ticket listing configs
	listingConfigs := conf.CombinedTicketListingConfigs()

	// Print config
	config.PrintTicketListingConfigs(listingConfigs)

	// Initial execution
	fetchAndProcessTickets(client, notificationClients, listingConfigs)

	// Create ticker
	ticker := time.NewTicker(refetchTime)
	defer ticker.Stop()

	// Loop until exit
	exitChan := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets(client, notificationClients, listingConfigs)
		case <-exitChan:
			return
		}
	}
}

func fetchAndProcessTickets(
	twicketsClient *twigots.Client,
	notificationClients map[config.NotificationType]notification.Client,
	listingConfigs []config.TicketListingConfig,
) {
	checkTime := time.Now()
	defer func() {
		lastCheckTime = checkTime
	}()

	fetchedListings, err := twicketsClient.FetchTicketListings(
		context.Background(),
		twigots.FetchTicketListingsInput{
			// Required
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
	if len(fetchedListings) == maxNumTickets {
		slog.Warn("Fetched the max number of tickets allowed. It is possible tickets have been missed.")
	}

	filteredListings := filterTicketListings(fetchedListings, listingConfigs)
	for _, matchedListing := range filteredListings {
		listing := matchedListing.listing
		listingConfig := matchedListing.config

		slog.Info(
			"found tickets for a wanted event",
			"wantedEventName", listingConfig.Event,
			"matchedEventName", listing.Event.Name,
			"numTickets", listing.NumTickets,
			"ticketPrice", listing.TotalPriceInclFee().String(),
			"originalTicketPrice", listing.OriginalTicketPrice().String(),
			"link", listing.URL(),
		)

		for _, notificationType := range listingConfig.Notification {
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

type matchedListingAndConfig struct {
	listing twigots.TicketListing
	config  config.TicketListingConfig
}

func filterTicketListings(
	listings twigots.TicketListings,
	listingConfigs []config.TicketListingConfig,
) []matchedListingAndConfig {
	matchedListings := make([]matchedListingAndConfig, 0, len(listings))
	for _, listing := range listings {
		for _, listingConfig := range listingConfigs {
			if ticketListingMatchesConfig(listing, listingConfig) {
				matchedListing := matchedListingAndConfig{
					config:  listingConfig,
					listing: listing,
				}
				matchedListings = append(matchedListings, matchedListing)
			}
		}
	}
	return matchedListings
}

func ticketListingMatchesConfig(listing twigots.TicketListing, listingConfig config.TicketListingConfig) bool {
	// Check name
	checkName := filter.EventName(listingConfig.Event, *listingConfig.EventSimilarity)
	if !checkName(listing) {
		return false
	}

	// Check regions
	checkRegions := filter.EventRegion(listingConfig.Regions...)
	if !checkRegions(listing) {

		wantedRegionStrings := make([]string, 0, len(listingConfig.Regions))
		for _, region := range listingConfig.Regions {
			wantedRegionStrings = append(wantedRegionStrings, region.Value)
		}
		wantedRegions := strings.Join(wantedRegionStrings, ", ")

		slog.Warn(
			"found tickets for a wanted event, but region does not match one of the regions wanted",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedRegions", wantedRegions,
			"listingRegion", listing.Event.Venue.Location.Region,
		)
		return false
	}

	// Check number of tickets
	numTickets := lo.FromPtr(listingConfig.NumTickets)
	checkNumTickets := filter.NumTickets(numTickets)
	if !checkNumTickets(listing) {
		slog.Warn(
			"found tickets for a wanted event, but number of tickets does not match the number wanted",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedNumTickets", listingConfig.NumTickets,
			"listingNumTickets", listing.NumTickets,
		)
		return false
	}

	// Check discount
	// If discount is close to 0 (e.g. 0 or a floating point error), set to -1 to allow any discount
	// Otherwise divide by 100 to get a number between 0-1
	discount := lo.FromPtr(listingConfig.Discount)
	if math.Abs(discount) < 1e-5 {
		discount = -1
	} else {
		discount /= 100
	}
	checkDiscount := filter.MinDiscount(discount)
	if !checkDiscount(listing) {
		slog.Warn(
			"found tickets for a wanted event, but discount does not match the discount wanted",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedDiscount", listingConfig.Discount,
			"listingDiscount", listing.Discount(),
		)
		return false
	}

	return true
}
