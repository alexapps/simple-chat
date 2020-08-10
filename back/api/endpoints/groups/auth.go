package auth

import (
	logInst "github.com/alexapps/simple-chat/back/api/log"

	"github.com/labstack/echo"
)

func AuthGroup(e *echo.Echo, log *logInst.Logger) {
	ah := auth.NewAuthHandler(log)
	e.POST("/signup", ah.SignUp)
	e.POST("/signin", ah.SignIn)
	e.POST("/signout", ah.SignOut)
}
