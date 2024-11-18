package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (InDepthLocation, error) {
	url := baseUrl + "/location-area"
	url = fmt.Sprintf("%s/%s", url, locationName)

	if val, ok := c.cache.Get(locationName); ok {
		locationInDepth := InDepthLocation{}
		err := json.Unmarshal(val, &locationInDepth)
		if err != nil {
			return InDepthLocation{}, err
		}
		return locationInDepth, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return InDepthLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return InDepthLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return InDepthLocation{}, err
	}

	locationInDepth := InDepthLocation{}
	err = json.Unmarshal(data, &locationInDepth)
	if err != nil {
		return InDepthLocation{}, err
	}
	c.cache.Add(locationName, data)

	return locationInDepth, nil
}
