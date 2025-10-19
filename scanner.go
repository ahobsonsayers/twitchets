package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twigots/filter"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/samber/lo"
)

func NewTicketsScanner(tsc TicketsScannerConfig) *TicketsScanner {
	return &TicketsScanner{
		config: tsc,
	}
}

type TicketsScannerConfig struct {
	twicketsClient      *twigots.Client
	notificationClients map[config.NotificationType]notification.Client
	listingConfigs      []config.TicketListingConfig
	refetchTime         time.Duration
}

type TicketsScanner struct {
	config TicketsScannerConfig

	// Synchronisation
	running      atomic.Bool
	runningWg    sync.WaitGroup
	shutdownOnce sync.Once

	// Cancellation
	cancel context.CancelFunc
}

func (w *TicketsScanner) IsRunning() bool   { return w.running.Load() }
func (w *TicketsScanner) WaitUntilStopped() { w.runningWg.Wait() }

// Start the scanner - blocks until stopped
func (w *TicketsScanner) Start(ctx context.Context) error {
	swapped := w.running.CompareAndSwap(false, true)
	if !swapped {
		return errors.New("worker already running")
	}
	w.runningWg.Add(1)
	defer w.cleanup()

	// Initial ticket scan
	fetchAndProcessTickets(w.config)

	// Create ticker
	ticker := time.NewTicker(w.config.refetchTime)
	defer ticker.Stop()

	// Create cancellable context
	ctx, cancel := context.WithCancel(ctx)
	w.cancel = cancel

	// Start ticket scanning
	for {
		select {
		case <-ticker.C:
			fetchAndProcessTickets(w.config)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *TicketsScanner) cleanup() {
	w.shutdownOnce.Do(func() {
		// Signal we are no longer running
		w.running.Store(false)
		w.runningWg.Done()
	})
}

// Stop the worker - blocks until stopped
func (w *TicketsScanner) Stop() {
	if w.cancel != nil {
		w.cancel()
	}
	w.WaitUntilStopped()
}

func fetchAndProcessTickets(tsc TicketsScannerConfig) {
	numTickets := maxNumTickets
	if latestTicketTime.IsZero() {
		numTickets = 10
	}

	// Fetch tickets listings from the twickets live feed
	fetchedListings, err := tsc.twicketsClient.FetchTicketListings(
		context.Background(),
		twigots.FetchTicketListingsInput{
			// Required
			Country: twigots.CountryUnitedKingdom,
			// Optional
			CreatedBefore: time.Now(),
			CreatedAfter:  latestTicketTime,
			MaxNumber:     numTickets,
		},
	)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Debug("Fetched tickets.", "numNewTickets", len(fetchedListings))

	if len(fetchedListings) == 0 {
		return
	}

	if len(fetchedListings) == maxNumTickets {
		slog.Warn("Fetched the max number of tickets allowed per check. It is possible tickets have been missed.")
	}

	// Update latest ticket time. Most recent ticket is first
	latestTicketTime = fetchedListings[0].CreatedAt.Time

	// Filter fetched ticket listings to those wanted
	filteredListings := filterTicketListings(fetchedListings, tsc.listingConfigs)
	for _, matchedListing := range filteredListings {

		listing := matchedListing.listing
		listingConfig := matchedListing.config

		// Log info about found ticket listing
		slog.Info(
			"Found tickets for a wanted event.",
			"wantedEventName", listingConfig.Event,
			"matchedEventName", listing.Event.Name,
			"numTickets", listing.NumTickets,
			"ticketPrice", listing.TotalPriceInclFee().String(),
			"originalTicketPrice", listing.OriginalTicketPrice().String(),
			"link", listing.URL(),
			"timeListed", listing.CreatedAt.Local(),
		)

		// Send notifications
		for _, notificationType := range listingConfig.Notification {

			notificationClient, ok := tsc.notificationClients[notificationType]
			if !ok {
				continue
			}

			err := notificationClient.SendTicketNotification(listing)
			if err != nil {
				slog.Error(
					"Failed to send notification.",
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
			"Found tickets for a wanted event, but region is not in allowed list.",
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
			"Found tickets for a wanted event, but number of tickets is incorrect.",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedNumTickets", numTickets,
			"listingNumTickets", listing.NumTickets,
		)
		return false
	}

	// Check discount
	// If value is close to 0 (e.g. 0 or a floating point error), set to -1 to allow any discount
	// Otherwise divide by 100 to get a number between 0-1
	discount := changeZeroToNegative(lo.FromPtr(listingConfig.MinDiscount)) / 100
	checkDiscount := filter.MinDiscount(discount)
	if !checkDiscount(listing) {
		slog.Warn(
			"Found tickets for a wanted event, but discount is too low.",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedDiscount", fmt.Sprintf("%.2f", discount),
			"listingDiscount", fmt.Sprintf("%.2f", listing.Discount()),
		)
		return false
	}

	// Check max ticket price including fee
	// If value is close to 0 (e.g. 0 or a floating point error), set to -1 to allow any discount
	// Otherwise divide by 100 to get a number between 0-1
	price := changeZeroToNegative(lo.FromPtr(listingConfig.MaxTicketPriceInclFee))
	checkMaxTicketPriceInclFee := filter.MaxTicketPriceInclFee(price)
	if !checkMaxTicketPriceInclFee(listing) {
		slog.Warn(
			"Found tickets for a wanted event, but ticket price including fee is too high.",
			"wantedEvent", listingConfig.Event,
			"listingEvent", listing.Event.Name,
			"wantedPrice", fmt.Sprintf("£%.2f", price),
			"listingPrice", listing.TicketPriceInclFee().String(),
		)
		return false
	}

	return true
}

// changeZeroToNegative changes a value that is 0 (or close to zero e.g. floating point error)
// to a negative number - specifically -1
func changeZeroToNegative(value float64) float64 {
	if math.Abs(value) < 1e-5 {
		return -1.0
	}
	return value
}
