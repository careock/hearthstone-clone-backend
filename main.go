// main.go
package main

import (
	"log"
	"net/http"

	"hearthstone-clone-backend/handlers"
	"hearthstone-clone-backend/utils"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define endpoints
	r.HandleFunc("/create-room", handlers.CreateRoom).Methods("POST")
	r.HandleFunc("/join-game", handlers.JoinGame).Methods("POST")
	r.HandleFunc("/ws", handlers.HandleWebSocket) // Add WebSocket endpoint

	// Start the hub
	go utils.StartHub()

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
