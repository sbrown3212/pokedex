package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := map[string]struct {
		key string
		val []byte
	}{
		"example": {
			key: "https://example.com",
			val: []byte("testdata"),
		},
		"example_path": {
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(tc.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
