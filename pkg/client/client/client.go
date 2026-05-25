package client

import (
	"bufio"
	"net"
)

type Client struct {
	Conn   net.Conn
	Reader *bufio.Reader
}
