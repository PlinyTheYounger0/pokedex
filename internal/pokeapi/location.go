package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocation(locationName string) (ResDeepLocation, error) {
	url := baseURL + "/location-area/" + locationName

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResDeepLocation{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ResDeepLocation{}, err
		}
		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return ResDeepLocation{}, err
		}

		c.cache.Add(url, val)
	}

	locationRes := ResDeepLocation{}
	if err := json.Unmarshal(val, &locationRes); err != nil {
		return ResDeepLocation{}, err
	}

	return locationRes, nil
}
