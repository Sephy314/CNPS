package main

import (
	"github.com/Sephy314/cnps/pkg/client/client"
	"github.com/Sephy314/cnps/pkg/dto"
	"github.com/Sephy314/cnps/pkg/logger"
)

func main() {
	conn, err := client.Connect(":31415")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	res, err := conn.Request(dto.Request{
		Target:  ":31415",
		Type:    "REQ",
		Cmd:     ".test",
		Act:     "action",
		Info:    dto.Info{},
		Payload: nil,
	})

	if err != nil {
		panic(err)
	}

	logger.Log{
		Msg:   res,
		Level: logger.INFO,
	}.Print()

}
