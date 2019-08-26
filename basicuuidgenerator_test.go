package basicuuidgenerator

import (
	"testing"
)

func TestNewV4(t *testing.T) {
	uuids := make(map[string]int)

	for i := 0; i < 100; i++ {
		uuid, err := NewV4()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		uuids[uuid] += 1
	}

	for _, count := range uuids {
		if count != 1 {
			t.Errorf("UUID not uniq")
		}
	}
}

func BenchmarkNewV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV4()
	}
}
