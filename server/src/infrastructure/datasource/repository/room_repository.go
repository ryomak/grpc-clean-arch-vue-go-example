package repository

import (
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
	"map-friend/src/infrastructure/datasource/sql_handler"
)

type roomRepository struct {
	sql_handler.ISqlHandler
}

func (r *roomRepository) GetRoomByName(name string) *model.Room {
	return &model.Room{
		Name: "test",
	}
}

func NewIRoomRepository(handler sql_handler.ISqlHandler) repository.IRoomRepository {
	return &roomRepository{handler}
}
