package server

import (
	"context"
	"errors"

	"map-friend/src/application/usecase"
	"map-friend/src/infrastructure/datasource/database"
	"map-friend/src/infrastructure/datasource/repository"
	"map-friend/src/interface/rpc"
	"map-friend/src/interface/rpc/transform"
)

func NewGrpcRoomServer(dbm database.DBM) *grpcRoomServer {
	usecase := usecase.NewIRoomUseCase(
		repository.NewITransactionRepository(dbm),
		repository.NewIRoomRepository(dbm),
		repository.NewIUserRepository(dbm),
	)
	return &grpcRoomServer{usecase}
}

type grpcRoomServer struct {
	Usecase usecase.IRoomUseCase
}

func (s *grpcRoomServer) GetRoom(ctx context.Context, req *rpc.RoomRequest) (*rpc.RoomResponse, error) {
	if req == nil {
		return nil, errors.New("not selected room")
	}
	mRoom := transform.TransformRoomModel(req)
	room, err := s.Usecase.GetRoom(ctx, mRoom.Name, mRoom.Users[0])
	if err != nil {
		return nil, err
	}
	return transform.TransformRoomRpc(room), nil
}

func (s *grpcRoomServer) GetRoomStream(server rpc.RoomHandler_GetRoomStreamServer) error {
	return nil
}
