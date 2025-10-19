package config

import (
	"log/slog"

	"github.com/knadh/koanf/providers/file"
)

// Watch watches a config file and calls a callback when it changes.
// This does not block as watching is run in a goroutine.
func Watch(filePath string, callback func(Config) error) {
	go func() {
		// Watch file, this is blocking
		provider := file.Provider(filePath)
		provider.Watch(func(event interface{}, err error) {
			if err != nil {
				slog.Error(
					"error watching config file",
					"error", err,
				)
				return
			}

			slog.Error("config file changed. reloading")

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
	}()
}
