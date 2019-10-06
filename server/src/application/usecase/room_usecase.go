package usecase

import (
	"context"
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
)

type IRoomUseCase interface {
	GetRoom(context.Context, string, *model.User) (*model.Room, error)
}

type roomUseCase struct {
	transaction repository.ITransactionRepository
	roomRepo    repository.IRoomRepository
	userRepo    repository.IUserRepository
}

func NewIRoomUseCase(
	t repository.ITransactionRepository,
	room repository.IRoomRepository,
	user repository.IUserRepository) IRoomUseCase {
	return &roomUseCase{t, room, user}
}

func (r *roomUseCase) GetRoom(ctx context.Context, roomName string, user *model.User) (*model.Room, error) {

	ctx, err := r.transaction.Begin(ctx)
	if err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}

	room, err := r.roomRepo.Save(ctx, roomName)
	if err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}

	//自身のデータ更新
	user.RoomName = roomName
	r.userRepo.Save(ctx, user)

	room.Users, err = r.userRepo.GetListByRoom(ctx, roomName)
	if err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}

	if _, err := r.transaction.Commit(ctx); err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}
	return room, nil
}

/*
func (r *roomRepository) GetRoom(roomName string, user *model.User) (*model.Room, error) {
	room := new(model.Room)
	tx, _ := r.DB.Begin()
	row := tx.QueryRow("select * from rooms where name = ?", roomName)
	//roomなければ作成
	if err := row.Scan(
		&room.Name,
		&room.Created,
		&room.Updated,
	); err != nil {
		stmt := `insert into rooms (name) values(?)`
		_, err := tx.Exec(stmt, roomName)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		room.Created = time.Now()
	}
	//userの情報を更新
	stmt := `insert into users (room_name,user_name,latitude,longitude) values (?,?,?,?)
			on DUPLICATE key update latitude = ? , longitude = ?`
	if _, err := tx.Exec(stmt,
		roomName, user.Name, user.Coordinate.Latitude, user.Coordinate.Longitude,
		user.Coordinate.Latitude, user.Coordinate.Longitude,
	); err != nil {
		tx.Rollback()
		return nil, err
	}

	rows, err := tx.Query("select * from users where room_name = ?", roomName)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	users := make([]model.User, 0, 16)
	for rows.Next() {
		var s model.User
		s.Coordinate = new(model.Coordinate)
		if err = rows.Scan(
			&s.Name,
			&s.RoomName,
			&s.Coordinate.Latitude,
			&s.Coordinate.Longitude,
			&s.Created,
			&s.Updated,
		); err != nil {
			tx.Rollback()
			return nil, err
		}
		users = append(users, s)
	}
	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &model.Room{
		Name:    roomName,
		Users:   users,
		Created: room.Created,
	}, tx.Commit()
}
*/
