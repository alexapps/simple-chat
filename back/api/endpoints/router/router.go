package router

import (
	"bitbucket.org/and_and/nbb-go-api/endpoints/groups"
	"bitbucket.org/and_and/nbbnub/logger"
	"github.com/labstack/echo"
)

func New(log *logger.Logger) *echo.Echo {
	e := echo.New()
	groups.AuthGroup(e, log)
	groups.ChatGroup(e, log)

	return e
}
