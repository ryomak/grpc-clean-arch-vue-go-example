package server

import (
	"context"

	"map-friend/src/application/usecase"
	"map-friend/src/infrastructure/datasource/repository"
	"map-friend/src/infrastructure/datasource/sql_handler"
	"map-friend/src/interface/rpc"
	"map-friend/src/interface/rpc/transform"
)

func NewGrpcRoomServer(sqlHandler sql_handler.ISqlHandler) *grpcRoomServer {
	usecase := usecase.NewIRoomUseCase(repository.NewIRoomRepository(sqlHandler))
	return &grpcRoomServer{usecase}
}

type grpcRoomServer struct {
	usecase.IRoomUseCase
}

func (s *grpcRoomServer) GetRoom(ctx context.Context, req *rpc.RoomName) (*rpc.Room, error) {
	return transform.TransformRoomRpc(s.GetRoomData(req.GetName())), nil
}
