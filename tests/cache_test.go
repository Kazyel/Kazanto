package tests

import (
	"testing"
	"time"

	"github.com/Kazyel/Poke-CLI/cache"
)

func TestCache(t *testing.T) {
	cache := cache.NewCache(0)

	cache.AddToCache("test", []byte("test"))

	val, ok := cache.GetFromCache("test")

	t.Logf("value")

	if !ok {
		t.Errorf("Expected to find value in cache.")
	}

	if string(val) != "test" {
		t.Errorf("Expected value to be 'test', got '%s'.", val)
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := cache.NewCache(baseTime)

	cache.AddToCache("https://example.com", []byte("testdata"))

	_, ok := cache.GetFromCache("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.GetFromCache("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
