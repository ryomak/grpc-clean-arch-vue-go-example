package usecase

import "github.com/ryomak/map-friend/server/src/domain/repository"

type IRoomUseCase interface {
}

type roomUseCase struct {
	repository.IRoomRepository
}

func NewRoomUseCase(r repository.IRoomRepository) IRoomUseCase {
	return &roomUseCase{r}
}
