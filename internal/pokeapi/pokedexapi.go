package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokedexcache"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokedexcache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache:      pokedexcache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: time.Minute},
	}

}
