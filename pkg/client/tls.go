package client

import (
	"bufio"
	"crypto/tls"
)

type TLSModule struct {
	config *tls.Config
}

func NewTLSClient(
	addr string,
	cfg *tls.Config,
) (*Client, error) {

	conn, err := tls.Dial(
		"tcp",
		addr,
		cfg,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn:   conn,
		Reader: bufio.NewReader(conn),
	}, nil
}
