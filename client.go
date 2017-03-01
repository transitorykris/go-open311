package open311

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client is used for communicating with an Open311 API
type Client struct {
	jurisdiction string
	url          string
	key          string
}

// New creates a brand new Client
func New(key string, jurisdiction string, apiURL string) *Client {
	c := &Client{
		jurisdiction: jurisdiction,
		url:          apiURL,
		key:          key,
	}
	return c
}

// perform an HTTP GET with the necessary goodies
func (c *Client) get(path string, v interface{}) error {
	resp, err := http.Get(c.url + path)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v); err != nil {
		return err
	}
	return nil
}
