package client

import (
	"github.com/asvinicius/actnsgo/internal/response"
)

func (c *Client) GetStatus() (*response.StatusResponse, error) {

	var data response.StatusResponse

	if err := c.get("/mercado/status", &data); err != nil {
		return nil, err
	}

	return &data, nil
}
