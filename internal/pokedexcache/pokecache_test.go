package pokedexcache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache()
	cache.Add("keyA", []byte("valA"))
	actual, ok := cache.Get("keyA")
	if !ok {
		t.Error("keyA does not exist")
	}
	if string(actual) != "valA" {
		t.Error("valA doesnt match")
	}
}
