package groupchat

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lolmourne/go-groupchat/model"
)

type RedisResource struct {
	rdb  *redis.Client
	next DBItf
}

type DBResource struct {
	db *sqlx.DB
}

type DBItf interface {
	GetChats(roomID int64) ([]model.Chat, error)
}

func NewRedisResource(rdb *redis.Client, next DBItf) DBItf {
	return &RedisResource{
		rdb:  rdb,
		next: next,
	}
}

func NewDBResource(dbParam *sqlx.DB) DBItf {
	return &DBResource{
		db: dbParam,
	}
}
