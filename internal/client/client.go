package client

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	url  string
	http *http.Client
}

func NewClient(url string) *Client {

	return &Client{
		url:  url,
		http: &http.Client{Timeout: 5 * time.Second},
	}

}

func (c *Client) get(path string, target any) error {

	resp, err := c.http.Get(c.url + path)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
