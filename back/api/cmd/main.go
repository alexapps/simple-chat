package main

import (
	"github.com/alexapps/simple-chat/back/api/config"
	constants "github.com/alexapps/simple-chat/back/api/constants"
	"github.com/alexapps/simple-chat/back/api/endpoints/router"
	logInst "github.com/alexapps/simple-chat/back/api/log"
)

func main() {
	// Init tools
	logger := &logInst.Logger{}
	conf, err := config.InitConfiguration()
	if err != nil {
		logger.Log(constants.LOG_ERROR, "Read config")
	}

	logger.Log(constants.LOG_INFO, "API service is started")
	e := router.New(logger)
	err = e.Start(":" + conf.HTTP.Port)
	if err != nil {
		logger.Log(constants.LOG_FATAL, "Echo server")
	}

}
