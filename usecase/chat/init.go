package groupchat

import (
	"github.com/lolmourne/go-groupchat/model"
	chat "github.com/lolmourne/go-groupchat/resource/chat"
)

type UseCase struct {
	dbRoomRsc chat.DBItf
}

type UsecaseItf interface {
	GetChats(roomID int64) ([]model.Chat, error)
}

func NewUseCase(dbRsc chat.DBItf) UsecaseItf {
	return UseCase{
		dbRoomRsc: dbRsc,
	}

}
