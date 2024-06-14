package main

import (
	"fmt"

	"github.com/Luccas71/gopportunities.git/config"
	"github.com/Luccas71/gopportunities.git/router"
)

func main() {
	//initialize configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// initialize router
	router.Initialize()
}
