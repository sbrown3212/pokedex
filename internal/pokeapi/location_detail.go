package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationDetail(area string) (RespDetailLocation, error) {
	url := baseURL + "/location-area" + "/" + area

	val, ok := c.cache.Get(url)
	if ok {
		var cachedLocationDetail RespDetailLocation
		err := json.Unmarshal(val, &cachedLocationDetail)
		if err != nil {
			return RespDetailLocation{}, err
		}

		return cachedLocationDetail, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RespDetailLocation{}, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailLocation{}, err
	}

	c.cache.Add(url, data)

	var locationDetail RespDetailLocation
	err = json.Unmarshal(data, &locationDetail)
	if err != nil {
		return RespDetailLocation{}, err
	}

	return locationDetail, nil
}
