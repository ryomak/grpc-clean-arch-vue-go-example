package usecase

import (
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
)

type IRoomUseCase interface {
	GetRoom(string, *model.User) (*model.Room, error)
}

type roomUseCase struct {
	Repo repository.IRoomRepository
}

func NewIRoomUseCase(r repository.IRoomRepository) IRoomUseCase {
	return &roomUseCase{r}
}

func (r *roomUseCase) GetRoom(roomName string, user *model.User) (*model.Room, error) {
	return r.Repo.GetRoom(roomName, user)
}
