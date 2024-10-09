package config

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"

	"github.com/knadh/koanf/parsers/yaml"
)

func Load(filePath string) (Config, error) {
	// Load config.
	k := koanf.New(".")
	err := k.Load(file.Provider(filePath), yaml.Parser())
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
