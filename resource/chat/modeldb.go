package groupchat

import (
	"database/sql"
	"time"
)

type ChatDB struct {
	ChatID    sql.NullInt64  `db:"db:chat_id"`
	RoomID    sql.NullInt64  `db:"db:room_id"`
	UserID    sql.NullInt64  `db:"db:user_id"`
	Message   sql.NullString `db:"db:message"`
	CreatedAt time.Time      `db:"created_at"`
}
