package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageUrl *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("CACHE HIT")
		// if we get a cache hit
		locationAreasResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreaResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("CACHE MISS")

	// make a request, if err return an empty slice
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// actually make a request, if err return an empty slice
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// check the statuscode, if err, you know what to do
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// get the data from res.body, if err... i dont need to repeat this
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreasResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullUrl, data)
	return locationAreasResp, nil
}
