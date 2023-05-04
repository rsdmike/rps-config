package main

import (
	"log"

	"github.com/rsdmike/rps/config"
	"github.com/rsdmike/rps/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
