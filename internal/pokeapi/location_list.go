package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResShallowLocations{}, fmt.Errorf("ListLocations Request Error: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ResShallowLocations{}, fmt.Errorf("")
		}
		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return ResShallowLocations{}, err
		}

		c.cache.Add(url, val)
	}

	locationsRes := ResShallowLocations{}
	if err := json.Unmarshal(val, &locationsRes); err != nil {
		return ResShallowLocations{}, err
	}

	return locationsRes, nil
}
