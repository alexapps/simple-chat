package models

type ChatMessage struct {
	ID       int64  `json:"ID,omitempty"`
	ChatID   int64  `json:"chatID,omitempty"`
	UserID   int64  `json:"userID,omitempty"`
	Body     string `json:"body,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Time     int64  `json:"time,omitempty"`
}

type Chat struct {
	ID             int64
	DateOfCreation int64
	Participants   []int64 `sql:",array" json:"participants"`
}
