package server

import (
	"log"
	"map-friend/src/infrastructure/datasource/sql_handler"
	"map-friend/src/interface/rpc"

	"net"

	"google.golang.org/grpc"
)

func GrpcRun() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcSrv := grpc.NewServer()
	rpc.RegisterRoomHandlerServer(grpcSrv, NewGrpcRoomServer(
		sql_handler.NewISqlHandler(),
	))
	log.Printf("grpc is running!")
	grpcSrv.Serve(listener)
}
