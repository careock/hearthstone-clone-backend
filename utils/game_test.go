package utils

import (
	"regexp"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
	// Генерируем ID
	id := GenerateRandomID()

	// Проверяем, что ID не пустой
	if id == "" {
		t.Error("GenerateRandomID() returned empty string")
	}

	// Проверяем формат UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(id) {
		t.Errorf("GenerateRandomID() returned invalid UUID format: %s", id)
	}

	// Проверяем уникальность - генерируем несколько ID и убеждаемся, что они разные
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
