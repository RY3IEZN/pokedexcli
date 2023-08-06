package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint

	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("CACHE HIT")
		// if we get a cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Println("CACHE MISS")

	// make a request, if err return an empty slice
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// actually make a request, if err return an empty slice
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	// check the statuscode, if err, you know what to do
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// get the data from res.body, if err... i dont need to repeat this
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)
	return pokemon, nil
}
