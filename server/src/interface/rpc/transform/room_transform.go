package transform

import (
	"map-friend/src/domain/model"
	"map-friend/src/interface/rpc"
)

func TransformRoomModel(rRoom *rpc.Room) *model.Room {
	users := make([]model.User, len(rRoom.GetUsers()))
	for i := 0; i < len(rRoom.GetUsers()); i++ {
		users[i] = *TransformUserModel(rRoom.GetUsers()[i])
	}
	return &model.Room{
		Name:  rRoom.GetName(),
		Users: users,
	}
}

func TransformRoomRpc(mRoom *model.Room) *rpc.Room {
	users := make([]*rpc.User, len(mRoom.Users))
	for i := 0; i < len(mRoom.Users); i++ {
		users[i] = TransformUserRpc(&mRoom.Users[i])
	}
	return &rpc.Room{
		Name:  mRoom.Name,
		Users: users,
	}
}
