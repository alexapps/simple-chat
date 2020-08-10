package models

type User struct {
	ID       int    `json:"ID"`
	Nickname string `json:"nickname"`
}

type SignInRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
