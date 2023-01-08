package utils

import (
	"testing"
)

func TestToMap(t *testing.T) {
	t.Run("Invalid json", func(t *testing.T) {
		mapData := ToMap("value")
		if len(mapData) > 0 {
			t.Error("Result should be empty")
		}
	})

	t.Run("Valid json but not map", func(t *testing.T) {
		mapData := ToMap("[1,2,3]")
		if len(mapData) > 0 {
			t.Error("Result should be empty")
		}
	})

	t.Run("Valid json", func(t *testing.T) {
		mapData := ToMap(map[string]interface{}{
			"name": "value",
		})
		if len(mapData) == 0 {
			t.Error("Result can not be empty")
		}
	})
}
