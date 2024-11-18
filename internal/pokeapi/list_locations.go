package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocation, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	var location RespShallowLocation
	err = json.Unmarshal(data, &location)
	if err != nil {
		return RespShallowLocation{}, err
	}
	c.cache.Add(url, data)

	return location, nil
}
