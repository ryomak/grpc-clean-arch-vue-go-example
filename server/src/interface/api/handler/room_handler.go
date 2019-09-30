package handler

import "net/http"

type Room struct {
	usecase.RoomUsecase
}

func GetRoomhandler() func(w http.ResponseWriter, r *http.Request) {
	repository := nil //repository.NewRepository():
	usecase := usecase.NewUsecase(repository)
	room := Room{usecase}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("aaaaa"))
	}
}
