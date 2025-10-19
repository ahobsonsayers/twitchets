package main

import (
	"context"
	"fmt"
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

	// Load user config
	userConfigPath := filepath.Join(cwd, "config.yaml")
	userConfig, err := config.Load(userConfigPath)
	if err != nil {
		log.Fatalf("config error:, %v", err)
	}

	// Get scanner config
	scannerConfig, err := ticketScannerConfigFromUserConfig(userConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tickets being scanned for
	config.PrintTicketListingConfigs(scannerConfig.ListingConfigs)

	// Create ticket scanner
	scanner := scanner.NewTicketScanner(scannerConfig)

	// Watch config file for changes (in a goroutine)
	go config.Watch(userConfigPath, func(conf config.Config) error {
		// Get scanner config
		scannerConfig, err := ticketScannerConfigFromUserConfig(userConfig)
		if err != nil {
			return err
		}

		// Update scanner config
		scanner.UpdateConfig(scannerConfig)

		return nil
	})

	// Start scanning for tickets
	slog.Info("Scanning for tickets...")
	err = scanner.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func ticketScannerConfigFromUserConfig(conf config.Config) (scanner.TicketScannerConfig, error) {
	// Create twickets client
	client, err := twigots.NewClient(conf.APIKey)
	if err != nil {
		return scanner.TicketScannerConfig{}, fmt.Errorf("failed to create twickets client: %w", err)
	}

	// Create notification clients
	notificationClients, err := notification.GetNotificationClients(conf.Notification)
	if err != nil {
		return scanner.TicketScannerConfig{}, fmt.Errorf("failed to create notification clients: %w", err)
	}

	// Get combined ticket listing configs
	listingConfigs := conf.CombinedTicketListingConfigs()

	return scanner.TicketScannerConfig{
		TwicketsClient:      client,
		NotificationClients: notificationClients,
		ListingConfigs:      listingConfigs,
		RefetchTime:         refetchTime,
	}, nil
}
