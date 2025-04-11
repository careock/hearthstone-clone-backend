// main.go

package main

import (
	"hearthstone-clone-backend/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/create-room", handlers.CreateRoom)
	http.HandleFunc("/join-game", handlers.JoinGame)

	http.ListenAndServe(":8080", nil)
}
