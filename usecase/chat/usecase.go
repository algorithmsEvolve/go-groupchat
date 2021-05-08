package groupchat

import (
	"github.com/lolmourne/go-groupchat/model"
)

func (u UseCase) GetChats(roomID int64) ([]model.Chat, error) {
	return u.dbRoomRsc.GetChats(roomID)
}
