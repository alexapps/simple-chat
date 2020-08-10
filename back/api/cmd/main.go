package main

import (
	c "github.com/alexapps/simple-chat/back/api/constants"
	l  "github.com/alexapps/simple-chat/back/api/log"
)

func main() {
	// Init tools
	logger := &Logger{}

	logger.Log(LOG_INFO, "API service is started")

}
