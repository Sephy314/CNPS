package client

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/types"
)

func (c *Client) Read() (*types.Response, error) {
	msg, err := c.Reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	var resp types.Response
	err = json.Unmarshal(msg, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}
