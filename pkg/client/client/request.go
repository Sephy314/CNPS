package client

import "github.com/Sephy314/cnps/pkg/dto"

func (c *Client) Request(r dto.Request) (*dto.Response, error) {
	err := c.Send(r)
	if err != nil {
		return nil, err
	}

	res, err := c.Read()

	if err != nil {
		return nil, err
	}

	return res, nil
}
