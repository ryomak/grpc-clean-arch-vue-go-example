package repository

import (
	"map-friend/src/domain/model"
)

type IRoomRepository interface {
	GetRoomByName(string) *model.Room
}
