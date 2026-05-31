package main

import (
	"crypto/tls"

	"github.com/Sephy314/cnps/pkg/client"
	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/types"
)

func main() {
	conn, err := client.NewTLSClient(":31415", &tls.Config{
		InsecureSkipVerify: true,
	})

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	res, err := conn.Request(types.Request{
		Target:  ":31415",
		Cmd:     ".test",
		Act:     types.QUERY,
		Info:    types.Info{},
		Payload: nil,
	})

	if err != nil {
		logger.Log{
			Msg:    "Error during request",
			Level:  logger.ERROR,
			Fields: err.Error(),
		}.Print()

		panic(err)
	}

	logger.Log{
		Msg:   res,
		Level: logger.INFO,
	}.Print()

}
