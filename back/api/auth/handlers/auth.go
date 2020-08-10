package auth

import (
	"net/http"

	log "github.com/alexapps/simple-chat/back/api/log"
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
func (ah *AuthHandler) SignUp() error {
	var rsp SignUpResponse
	//TODO
	return c.JSON(http.StatusOK, rsp)
}

// SignIn - log in
func (ah *AuthHandler) SignIn() error {
	//TODO
	return c.JSON(http.StatusOK, rsp)

}
func (ah *AuthHandler) SignOut() error {
	//TODO
	return c.JSON(http.StatusOK, rsp)

}
