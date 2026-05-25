package client

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/dto"
)

func (c *Client) Read() (*dto.Response, error) {
	msg, err := c.Reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	var resp dto.Response
	err = json.Unmarshal(msg, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}
