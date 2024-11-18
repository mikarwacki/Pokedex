package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(pokemonName); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(pokemon.Name, data)
	return pokemon, nil
}
