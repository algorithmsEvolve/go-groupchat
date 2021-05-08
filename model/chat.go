package model

import "time"

type Chat struct {
	ChatID    int64     `json:"chat_id"`
	RoomID    int64     `json:"room_id"`
	UserID    int64     `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
