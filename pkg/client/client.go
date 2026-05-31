package client

import (
	"bufio"
	"net"
)

type Client struct {
	Conn net.Conn
	//dial   func(string) (net.Conn, error)
	Reader *bufio.Reader
}
