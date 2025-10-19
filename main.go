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
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/scanner"
	"github.com/joho/godotenv"
)

//go:generate go tool oapi-codegen -config ./oapi.models.yaml ./schema/models.openapi.yaml
//go:generate go tool oapi-codegen -config ./oapi.server.yaml ./schema/server.openapi.yaml

const refetchTime = 1 * time.Minute

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
	notificationClients, err := notification.GetNotificationClients(conf.Notification)
	if err != nil {
		log.Fatal(err)
	}

	// Get combined ticket listing configs
	listingConfigs := conf.CombinedTicketListingConfigs()

	// Print config
	config.PrintTicketListingConfigs(listingConfigs)

	slog.Info("Scanning Tickets...")

	scanner := scanner.NewTicketScanner(scanner.TicketScannerConfig{
		TwicketsClient:      client,
		NotificationClients: notificationClients,
		ListingConfigs:      listingConfigs,
		RefetchTime:         refetchTime,
	})

	err = scanner.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
