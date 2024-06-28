package main

import (
	"github.com/hermangoncalves/go_opportunities/config"
	"github.com/hermangoncalves/go_opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
	}
	router.Initialize()
}
