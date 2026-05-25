package client

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/dto"
)

func (c *Client) Send(r dto.Request) error {
	marshaled, err := json.Marshal(r)
	if err != nil {
		return err
	}

	bytes := append(marshaled, '\n')

	_, err = c.Conn.Write(bytes)

	return err
}
