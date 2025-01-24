package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize configs
	err := config.InitDB()
	if err != nil {
		logger.Err("config initialization error")
		return
	}

	// Initalize routes
	router.InitRouter()
}
