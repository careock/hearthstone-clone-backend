package service

import (
	"regexp"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
	// Generate ID
	id := GenerateRandomID()

	// Check that ID is not empty
	if id == "" {
		t.Error("GenerateRandomID() returned empty string")
	}

	// Check UUID format
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(id) {
		t.Errorf("GenerateRandomID() returned invalid UUID format: %s", id)
	}

	// Check uniqueness
	ids := make(map[string]bool)
	ids[id] = true

	for i := 0; i < 100; i++ {
		newID := GenerateRandomID()
		if ids[newID] {
			t.Errorf("GenerateRandomID() generated duplicate ID: %s", newID)
		}
		ids[newID] = true
	}
}
