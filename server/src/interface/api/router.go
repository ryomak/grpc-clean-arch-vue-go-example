package api

import (
	"map-friend/src/interface/api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type httpMethod string

const (
	GetMethod    httpMethod = "GET"
	PostMethod   httpMethod = "POST"
	PutMethod    httpMethod = "PUT"
	DeleteMethod httpMethod = "DELETE"
)

type Route struct {
	Path    string
	Method  httpMethod
	Handler func(w http.ResponseWriter, r *http.Request)
}

func routes() []Route {
	return []Route{
		Route{
			Path:    "/",
			Method:  GetMethod,
			Handler: handler.GetRoomhandler(),
		},
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, v := range routes() {
		r.HandleFunc(v.Path, v.Handler).Methods(string(v.Method))
	}
	return r
}
