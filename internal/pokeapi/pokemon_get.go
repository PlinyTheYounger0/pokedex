package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, val)
	}

	pokemonRes := Pokemon{}
	if err := json.Unmarshal(val, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	return pokemonRes, nil
}
