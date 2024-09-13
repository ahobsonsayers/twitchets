package main

import (
	"errors"
	"fmt"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	Country twickets.Country  `koanf:"country"`
	Regions []twickets.Region `koanf:"regions"`
	Events  []EventConfig     `koanf:"events"`
}

func (c Config) EventNames() []string {
	eventNames := make([]string, 0, len(c.Events))
	for _, event := range c.Events {
		eventName := event.Name
		eventNames = append(eventNames, eventName)
	}
	return eventNames
}

func (c *Config) parseKoanf(k *koanf.Koanf) error {
	if k == nil {
		return nil
	}

	// Parse country
	countryString := k.String("country")
	if countryString == "" {
		return errors.New("country must be set")
	}
	country := twickets.Countries.Parse(countryString)
	if country == nil {
		return fmt.Errorf("%s is not a valid country code", countryString)
	}

	// Parse regions
	regionStrings := k.Strings("regions")
	regions := make([]twickets.Region, 0, len(regionStrings))
	for _, regionString := range regionStrings {
		region := twickets.Regions.Parse(regionString)
		if region == nil {
			return fmt.Errorf("%s is not a valid region code", countryString)
		}
		regions = append(regions, *region)
	}

	// Parse event config
	var events []EventConfig
	err := k.Unmarshal("events", &events)
	if err != nil {
		return fmt.Errorf("invalid events: %w", err)
	}

	c.Country = *country
	c.Regions = regions
	c.Events = events

	return nil
}

type EventConfig struct {
	Name       string `koanf:"name"`
	NumTickets int    `koanf:"num_tickets"`
}

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var (
	k      = koanf.New(".")
	parser = yaml.Parser()
)

func LoadConfig(filePath string) (Config, error) {
	// Load config.
	err := k.Load(file.Provider(filePath), yaml.Parser())
	if err != nil {
		return Config{}, fmt.Errorf("error loading config: %w", err)
	}

	// Parse config
	config := Config{}
	err = config.parseKoanf(k)
	if err != nil {
		return Config{}, fmt.Errorf("invalid config: %w", err)
	}

	return config, nil
}