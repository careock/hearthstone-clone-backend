package main

import (
	"hearthstone-clone-backend/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	if err := s.Start(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
