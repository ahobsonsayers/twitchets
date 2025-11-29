package config

import (
	"log/slog"

	"github.com/knadh/koanf/providers/file"
)

// Watch watches a config file and calling a callback when it changes.
// This function is blocking.
func Watch(filePath string, callback func(Config) error) error {
	provider := file.Provider(filePath)
	return provider.Watch(func(event interface{}, err error) {
		if err != nil {
			slog.Error(
				"error watching config file",
				"error", err,
			)
			return
		}

		slog.Info("config file changed. reloading")

		// Reload config
		config, err := Load(filePath)
		if err != nil {
			slog.Error(
				"error reloading config file",
				"error", err,
			)
			return
		}

		// Call callback
		err = callback(config)
		if err != nil {
			slog.Error(
				"error calling callback",
				"error", err,
			)
			return
		}
	})
}
