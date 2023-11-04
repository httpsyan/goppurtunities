package main

import (
	"github.com/cplxx/goppurtunities.git/config"
	"github.com/cplxx/goppurtunities.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
 
	err := config.Init()
	if err != nil {
		logger.ErrorF("Config initialization error: %v",err)
		return
	}

 router.Initialize()
}