package main

import (
	"errors"
	"fmt"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/knadh/koanf"
	yamlparser "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Config struct {
	APIKey        string            `json:"apiKey"`
	Country       twickets.Country  `json:"country"`
	GlobalConfig  GlobalEventConfig `json:"global"`
	TicketsConfig []TicketConfig    `json:"tickets"`
}

func (c Config) Validate() error {
	if c.APIKey == "" {
		return errors.New("api key must be set")
	}

	if c.Country.Value == "" {
		return errors.New("country must be set")
	}
	if !twickets.Countries.Contains(c.Country) {
		return fmt.Errorf("country '%s' is not valid", c.Country)
	}

	err := c.GlobalConfig.Validate()
	if err != nil {
		return fmt.Errorf("global config is not valid: %w", err)
	}

	for idx, ticketConfig := range c.TicketsConfig {
		err := ticketConfig.Validate()
		if err != nil {
			return fmt.Errorf("event config at index [%d] is no valid: %w", idx, err)
		}
	}

	return nil
}

func (c Config) Filters() []twickets.Filter {
	filters := make([]twickets.Filter, 0, len(c.TicketsConfig))
	for _, ticketConfig := range c.TicketsConfig {

		var filter twickets.Filter
		filter.Event = ticketConfig.Event.Name

		// Set name similarity
		if ticketConfig.Event.Similarity == nil {
			filter.EventSimilarity = c.GlobalConfig.EventSimilarity
		} else if *ticketConfig.Event.Similarity > 0 {
			filter.EventSimilarity = *ticketConfig.Event.Similarity
		}

		// Set regions
		if ticketConfig.Regions == nil {
			filter.Regions = c.GlobalConfig.Regions
		} else {
			filter.Regions = ticketConfig.Regions
		}

		// Set num tickets
		if ticketConfig.NumTickets == nil {
			filter.NumTickets = c.GlobalConfig.NumTickets
		} else if *ticketConfig.NumTickets > 0 {
			filter.NumTickets = *ticketConfig.NumTickets
		}

		// Set discount
		if ticketConfig.Discount == nil {
			filter.Discount = c.GlobalConfig.Discount
		} else if *ticketConfig.Discount > 0 {
			filter.Discount = *ticketConfig.Discount
		}

		filters = append(filters, filter)
	}

	return filters
}

// GlobalEventConfig is config that applies to all events,
// unless an event explicitly overwrites its.
// Country is required.
type GlobalEventConfig struct {
	EventSimilarity float64           `json:"eventSimilarity"`
	Regions         []twickets.Region `json:"regions"`
	NumTickets      int               `json:"numTickets"`
	Discount        float64           `json:"discount"`
}

func (c GlobalEventConfig) Validate() error {
	// Reuse the filter validation logic
	filter := twickets.Filter{
		Event:           "global", // Event must be be set - this is arbitrary
		EventSimilarity: c.EventSimilarity,
		Regions:         c.Regions,
		NumTickets:      c.NumTickets,
		Discount:        c.Discount,
	}
	err := filter.Validate()
	if err != nil {
		return err
	}

	return nil
}

type TicketConfig struct {
	Event           Event             `json:"event"`
	EventSimilarity *float64          `json:"eventSimilarity"`
	Regions         []twickets.Region `json:"regions"`
	NumTickets      *int              `json:"numTickets"`
	Discount        *float64          `json:"discount"`
}

func (t TicketConfig) Validate() error {
	if t.Event.Name == "" {
		return errors.New("event name must be set")
	}

	for _, region := range t.Regions {
		if !twickets.Regions.Contains(region) {
			return fmt.Errorf("region '%s' is not valid", region)
		}
	}

	return nil
}

type Event struct {
	Name       string   `json:"name"`
	Similarity *float64 `json:"Similarity"`
}

func (e *Event) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {

	case yaml.ScalarNode:
		var name string
		err := node.Decode(&name)
		if err != nil {
			return err
		}

		e.Name = name

	case yaml.MappingNode:
		type eventAlias Event
		var event eventAlias
		err := node.Decode(&event)
		if err != nil {
			return err
		}

		*e = Event(event)
	}

	return nil
}

func LoadConfig(filePath string) (Config, error) {
	// Load config.
	k := koanf.New(".")
	err := k.Load(file.Provider(filePath), yamlparser.Parser())
	if err != nil {
		return Config{}, fmt.Errorf("error loading config: %w", err)
	}

	// Parse config
	config, err := parseKoanf(k)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config: %w", err)
	}

	return config, nil
}

func parseKoanf(k *koanf.Koanf) (Config, error) {
	if k == nil {
		return Config{}, nil
	}

	// Parse config
	var config Config
	err := k.UnmarshalWithConf(
		"", nil,
		koanf.UnmarshalConf{
			Tag: "json",
			DecoderConfig: &mapstructure.DecoderConfig{
				// Mostly a copy of decoder config form koanf
				Result:           &config,
				WeaklyTypedInput: true,
				ErrorUnused:      true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToSliceHookFunc(","),
					mapstructure.TextUnmarshallerHookFunc()),
			},
		},
	)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	err = config.Validate()
	if err != nil {
		return Config{}, fmt.Errorf("invalid config: %w", err)
	}

	return config, nil
}
