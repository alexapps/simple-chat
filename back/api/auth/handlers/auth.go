package auth

import (
	"net/http"

	authModel "github.com/alexapps/simple-chat/back/api/auth/models"
	log "github.com/alexapps/simple-chat/back/api/log"
	"github.com/labstack/echo"
)

// AuthHandler -
type AuthHandler struct {
	logger *log.Logger
}

// NewAuthHandler -
func NewAuthHandler(loggerInst *log.Logger) *AuthHandler {
	return &AuthHandler{
		logger: loggerInst,
	}
}

// SignUp - register
func (ah *AuthHandler) SignUp(c echo.Context) error {
	var rsp authModel.SignUpResponse
	//TODO
	return c.JSON(http.StatusOK, rsp)
}

// SignIn - log in
func (ah *AuthHandler) SignIn(c echo.Context) error {
	var rsp authModel.SignInResponse
	//TODO
	return c.JSON(http.StatusOK, rsp)

}
func (ah *AuthHandler) SignOut(c echo.Context) error {
	var rsp authModel.SignOutResponse
	//TODO
	return c.JSON(http.StatusOK, rsp)

}
