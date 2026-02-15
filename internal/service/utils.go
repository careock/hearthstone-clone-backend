package service

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomID generates a random UUID string
func GenerateRandomID() string {
	return uuid.New().String()
}
