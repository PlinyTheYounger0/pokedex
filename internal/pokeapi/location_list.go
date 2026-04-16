package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	val, ok := c.cache.Get(*pageUrl)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResShallowLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ResShallowLocations{}, err
		}
		defer res.Body.Close()

		val, err := io.ReadAll(res.Body)
		if err != nil {
			return ResShallowLocations{}, err
		}

		c.cache.Add(*pageUrl, val)
	}

	locationsRes := ResShallowLocations{}
	if err := json.Unmarshal(val, &locationsRes); err != nil {
		return ResShallowLocations{}, err
	}

	return locationsRes, nil
}
