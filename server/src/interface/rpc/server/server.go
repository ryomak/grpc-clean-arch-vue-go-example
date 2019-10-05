package server

import (
	"fmt"
	"log"
	"map-friend/src/infrastructure/datasource/sql_handler"
	"map-friend/src/interface/rpc"

	"net"

	"google.golang.org/grpc"
)

func GrpcRun(dbConfigPath, env, addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcSrv := grpc.NewServer()
	rpc.RegisterRoomHandlerServer(grpcSrv, NewGrpcRoomServer(
		sql_handler.NewSqlHandler(dbConfigPath, env),
	))
	log.Printf(fmt.Sprintf("env:%s\nport%s\ngrpc server start... ", env, addr))
	grpcSrv.Serve(listener)
}
