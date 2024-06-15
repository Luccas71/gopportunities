package main

import (
	"github.com/Luccas71/gopportunities.git/config"
	"github.com/Luccas71/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	//logs
	logger := config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
