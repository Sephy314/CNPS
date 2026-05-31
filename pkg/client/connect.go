package client

import (
	"bufio"
	"net"
)

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, err
	}

	return &Client{
		Conn:   conn,
		Reader: bufio.NewReader(conn),
	}, nil
}
