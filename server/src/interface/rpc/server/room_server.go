package server

import (
	"context"
	"errors"

	"map-friend/src/application/usecase"
	"map-friend/src/infrastructure/datasource/repository"
	"map-friend/src/infrastructure/datasource/sql_handler"
	"map-friend/src/interface/rpc"
	"map-friend/src/interface/rpc/transform"
)

func NewGrpcRoomServer(sqlHandler *sql_handler.SqlHandler) *grpcRoomServer {
	usecase := usecase.NewIRoomUseCase(repository.NewIRoomRepository(sqlHandler))
	return &grpcRoomServer{usecase}
}

type grpcRoomServer struct {
	Usecase usecase.IRoomUseCase
}

func (s *grpcRoomServer) GetRoom(ctx context.Context, req *rpc.Room) (*rpc.Room, error) {
	if req == nil {
		return nil, errors.New("not selected room")
	}
	mRoom := transform.TransformRoomModel(req)
	room, err := s.Usecase.GetRoom(mRoom.Name, &mRoom.Users[0])
	if err != nil {
		return nil, err
	}
	return transform.TransformRoomRpc(room), nil
}
