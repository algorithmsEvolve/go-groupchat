package groupchat

import (
	"log"

	"github.com/lolmourne/go-groupchat/model"
)

func (dbr *DBResource) GetChats(roomID int64) ([]model.Chat, error) {
	query := `
		SELECT
			*
		FROM chat
		
		WHERE
			room_id = $1
	`
	chats, err := dbr.db.Queryx(query, roomID)

	var resultChats []model.Chat
	for chats.Next() {
		var r ChatDB
		err = chats.StructScan(&r)

		if err == nil {
			resultChats = append(resultChats, model.Chat{
				ChatID:    r.ChatID.Int64,
				RoomID:    r.RoomID.Int64,
				UserID:    r.UserID.Int64,
				Message:   r.Message.String,
				CreatedAt: r.CreatedAt,
			})
		}
	}
	log.Println("result chats:", resultChats)

	return resultChats, err
}
