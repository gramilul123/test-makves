package main

import (
	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/internal/app"
)

func main() {
	// Configuration
	cfg := config.New()

	// Run
	app.Run(cfg)
}
