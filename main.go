package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/config"
	"github.com/joho/godotenv"
)

const (
	maxNumTickets = 250
	refetchTime   = 1 * time.Minute
)

var latestTicketTime time.Time

func init() {
	_ = godotenv.Load()

	// Set debug level
	debugEnv := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugEnv)
	if err == nil && debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
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

	slog.Info("Scanning Tickets...")
	scanner := NewTicketsScanner(TicketsScannerConfig{
		twicketsClient:      client,
		notificationClients: notificationClients,
		listingConfigs:      listingConfigs,
		refetchTime:         refetchTime,
	})
	scanner.Start(context.Background())
}
