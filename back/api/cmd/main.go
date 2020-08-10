package main

import (
	"bitbucket.org/and_and/nbb-go-api/config"
	"bitbucket.org/and_and/nbb-go-api/router"
	constants "github.com/alexapps/simple-chat/back/api/constants"
	logInst "github.com/alexapps/simple-chat/back/api/log"
)

func main() {
	// Init tools
	logger := &logInst.Logger{}
	conf :=

		logger.Log(constants.LOG_INFO, "API service is started")

	e := router.New(&logger)
	err = e.Start(":" + config.HTTP.Port)
	if err != nil {
		logInst.Fatal(err.Error(), logger.Args{"source": "micro service"})
	}

}
