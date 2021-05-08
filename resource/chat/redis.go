package groupchat

import (
	"github.com/lolmourne/go-groupchat/model"
)

func (dbr *RedisResource) GetChats(roomID int64) ([]model.Chat, error) {
	//no need redis for DML (insert)
	return dbr.next.GetChats(roomID)
}
