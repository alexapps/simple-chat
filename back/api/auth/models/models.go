package models

/*
* Structs for auth
 */

// SignUpResponse  - Register
type SignUpResponse struct {
}

// SignInResponse - Log in
type SignInResponse struct {
	Token          string `json:"token"`
	TokenExpiredAt int64  `json:"tokenExpiredAt"`
}
