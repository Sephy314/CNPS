package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Sephy314/cnps/pkg/client"
	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/types"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s <server> <tls | tcp>", os.Args[0])
	}

	addr := os.Args[1]

	clientType := os.Args[2]

	var Conn *client.Client
	var err error

	if clientType == "tls" {
		Conn, err = client.NewTLSClient(addr, &tls.Config{
			InsecureSkipVerify: true,
		})

		if err != nil {
			panic(err)
		}

		defer Conn.Close()
	} else if clientType == "tcp" {
		Conn, err = client.NewClient(addr)

		if err != nil {
			panic(err)
		}

		defer Conn.Close()
	} else {
		logger.Log{
			Msg:    "Invalid client type",
			Level:  logger.ERROR,
			Fields: nil,
		}.Print()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("cnps %s >", addr)

		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		parts := strings.SplitN(line, " ", 2)
		cmd := parts[0]

		if len(parts) == 0 {
			continue
		}

		// refactor this code into a strategy pattern
		// I use switch cuz It's just MVP
		switch cmd {
		case "exit":
			return

		case "ping":
			fmt.Println("pong")

		case "req":
			if len(parts) != 2 {
				logger.Log{
					Msg:    "invalid args",
					Level:  logger.ERROR,
					Fields: nil,
				}.Print()
				continue
			}

			req := parts[1]

			fmt.Printf("Request to %v, Body : %v \n", addr, req)
			response, err := request(Conn, req)
			if err != nil {
				return
			}

			fmt.Printf("%+v\n", *response)

		default:
			fmt.Println("unknown command")
		}
	}
}

func request(c *client.Client, body string) (res *types.Response, e error) {
	var req types.Request
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		e = err
		return
	}

	response, err := c.Request(req)
	if err != nil {
		e = err
		return
	}

	res = response

	return
}
