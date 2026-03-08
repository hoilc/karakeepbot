// Package main is the entry point of the application. It initializes the
// configuration, sets up logging, and starts the Karakeepbot.
package main

import (
	"fmt"

	"github.com/Madh93/karakeepbot/internal/config"
	"github.com/Madh93/karakeepbot/internal/karakeepbot"
	"github.com/Madh93/karakeepbot/internal/logging"
)

// main initializes the configuration, sets up logging, and starts the
// Karakeepbot.
func main() {
	// Load configuration
	config := config.New()

	// Setup logger
	logger := logging.New(&config.Logging)
	if config.Path != "" {
		logger.Debug(fmt.Sprintf("Loaded configuration from %s", config.Path))
	}

	// Log list_id configuration for debugging
	logger.Debug(fmt.Sprintf("Karakeep List ID: %q", config.Karakeep.ListID))

	// Setup karakeepbot
	karakeepbot := karakeepbot.New(logger, config)

	// Let's go
	logger.Info("3, 2, 1... Launching Karakeepbot... 🚀")
	if err := karakeepbot.Run(); err != nil {
		logger.Fatal("💥 Something went wrong.", "error", err)
	}
}
