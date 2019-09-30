package main

import (
	"map-friend/src/interface/api"
	"net/http"
)

func main() {

	r := api.NewRouter()
	http.ListenAndServe(":8080", r)
}
