package model

type User struct {
	ID         uint
	Name       string
	RoomName   string
	Coordinate *Coordinate
}
