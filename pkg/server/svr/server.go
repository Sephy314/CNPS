package svr

import (
	"log"
	"net"

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
	log.SetFlags(0)

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

	log.Printf("CNPS pkg is listening on %v", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handler.HandleConnection(conn)
	}

}
