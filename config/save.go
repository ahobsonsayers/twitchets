package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

func Save(conf Config, filePath string) error {
	// Marshal config
	configBytes, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	// Write to file
	err = os.WriteFile(filePath, configBytes, 0o644)
	if err != nil {
		return fmt.Errorf("error writing config to file: %w", err)
	}

	return nil
}
