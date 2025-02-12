package groupchat

import (
	"github.com/lolmourne/go-groupchat/model"
)

func (u UseCase) CreateGroupchat(name string, adminID int64, desc string, categoryID int64) (model.Room, error) {

	err := u.dbRoomRsc.CreateRoom(name, adminID, desc, categoryID)
	if err != nil {
		return model.Room{}, err
	}

	return model.Room{
		Name:        name,
		AdminUserID: adminID,
		Description: desc,
		CategoryID:  categoryID,
	}, nil

}

func (u UseCase) EditGroupchat(name, desc, categoryID string) (model.Room, error) {
	panic("TBC")
}

func (u UseCase) JoinRoom(roomID, userID int64) error {
	err := u.dbRoomRsc.AddRoomParticipant(roomID, userID)

	return err
}

func (u UseCase) GetRoomByID(roomID int64) (model.Room, error) {
	return u.dbRoomRsc.GetRoomByID(roomID)
}

func (u UseCase) GetRoomList(userID int64) ([]model.Room, error) {
	return u.dbRoomRsc.GetRooms(userID)
}

func multiply(x, y int64) int64 {
	return x * y
}

func (u UseCase) CreateChat(roomID int64, userID int64, message string) (model.Chat, error) {

	err := u.dbRoomRsc.CreateChat(roomID, userID, message)

	if err != nil {
		return model.Chat{}, err
	}

	return model.Chat{
		RoomID:  roomID,
		UserID:  userID,
		Message: message,
	}, nil

}
