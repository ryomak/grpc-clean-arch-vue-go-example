package repository

import (
	"context"
	"map-friend/src/domain/repository"
	"map-friend/src/infrastructure/datasource/database"
	"time"

	"map-friend/src/domain/model"
)

func NewIRoomRepository(dbm database.DBM) repository.IRoomRepository {
	return &roomRepository{dbm}
}

type roomRepository struct {
	dbm database.DBM
}

func (r *roomRepository) Save(ctx context.Context, roomName string) (*model.Room, error) {
	query := `insert into rooms (name) values (?) on DUPLICATE key update updated = ?`
	stmt, err := r.dbm.Prepare(ctx, query)
	_, err = stmt.ExecContext(ctx, roomName, time.Now())
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, roomName)
}

func (r *roomRepository) Get(ctx context.Context, name string) (*model.Room, error) {
	room := &model.Room{}
	query := "select * from rooms where name = ?"
	stmt, err := r.dbm.Prepare(ctx, query)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRowContext(ctx, name).Scan(&room.Name, &room.Created, &room.Updated)
	if err != nil {
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
