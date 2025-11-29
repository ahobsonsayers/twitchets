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
	"github.com/ahobsonsayers/twitchets/server"
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
	ticketScannerConfig, err := ticketScannerConfigFromUserConfig(userConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tickets being scanned for
	config.PrintTicketListingConfigs(ticketScannerConfig.ListingConfigs)

	// Create ticket ticketScanner
	ticketScanner := scanner.NewTicketScanner(ticketScannerConfig)

	// Watch config file for changes (in a goroutine)
	go func() {
		err := config.Watch(
			userConfigPath,
			getUserConfigUpdateCallback(ticketScanner),
		)
		if err != nil {
			log.Fatalf("failed to set up config watching: %v", err)
		}
	}()

	// Start scanner in gorountine
	log.Println("Scanning for tickets...")
	go func() {
		err = ticketScanner.Start(context.Background())
		if err != nil {
			log.Fatalf("error running scanner: %v", err)
		}
	}()

	// Run server
	err = server.Start(9000, userConfigPath)
	if err != nil {
		log.Fatalf("error running server: %v", err)
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

func getUserConfigUpdateCallback(ticketScanner *scanner.TicketScanner) func(config.Config) error {
	return func(userConfig config.Config) error {
		// Get scanner config
		scannerConfig, err := ticketScannerConfigFromUserConfig(userConfig)
		if err != nil {
			return err
		}

		// Update scanner config
		ticketScanner.UpdateConfig(scannerConfig)

		return nil
	}
}
