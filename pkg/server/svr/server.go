package svr

import (
	"fmt"
	"net"

	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/handler"
)

func Start(addr string) error {

	err := startCNPSServer(addr)

	if err != nil {
		return err
	}

	return nil
}

func startCNPSServer(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			panic(err)
		}
	}(ln)

	logger.Log{
		Msg:   fmt.Sprintf("Starting server on %s", addr),
		Level: logger.INFO,
	}.Print()

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Log{
				Msg:   fmt.Sprintf("Error accepting connection: %v", err),
				Level: logger.ERROR,
			}.Print()

			continue
		}

		go handler.HandleConnection(conn)
	}

}
