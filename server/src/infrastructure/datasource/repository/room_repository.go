package repository

import (
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
	"map-friend/src/infrastructure/datasource/sql_handler"
	"time"
)

func NewIRoomRepository(handler *sql_handler.SqlHandler) repository.IRoomRepository {
	return &roomRepository{handler}
}

type roomRepository struct {
	*sql_handler.SqlHandler
}

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
