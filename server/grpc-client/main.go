package main

import (
	"context"
	"fmt"
	"os"

	"map-friend/src/interface/rpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}
	defer conn.Close()
	client := rpc.NewRoomHandlerClient(conn)
	req := &rpc.RoomName{Name: "test"}
	room, err := client.GetRoom(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "room: %v\n", room)
		fmt.Fprintf(os.Stderr, "room: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "user: %s\n", room.GetName())
}
