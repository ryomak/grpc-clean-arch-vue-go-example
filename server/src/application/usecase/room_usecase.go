package usecase

import (
	"fmt"
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
)

type IRoomUseCase interface {
	GetRoomData(string) *model.Room
}

type roomUseCase struct {
	repository.IRoomRepository
}

func NewIRoomUseCase(r repository.IRoomRepository) IRoomUseCase {
	return &roomUseCase{r}
}

func (r *roomUseCase) GetRoomData(roomName string) *model.Room {
	fmt.Println(r.GetRoomByName(roomName))
	return r.GetRoomByName(roomName)
}
