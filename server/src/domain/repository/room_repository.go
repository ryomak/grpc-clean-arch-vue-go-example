package repository

import (
	"map-friend/src/domain/model"
)

type IRoomRepository interface {
	GetRoom(string, *model.User) (*model.Room, error)
}
