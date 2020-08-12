package router

import (
	groups "github.com/alexapps/simple-chat/back/api/endpoints/groups"
	logger "github.com/alexapps/simple-chat/back/api/log"
	"github.com/labstack/echo"
)

func New(log *logger.Logger) *echo.Echo {
	e := echo.New()
	groups.AuthGroup(e, log)
	groups.ChatGroup(e, log)

	return e
}
