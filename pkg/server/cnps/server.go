package cnps

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/handler"
)

type Server struct {
	addr     string
	listener net.Listener
	wrapConn func(net.Conn) net.Conn
}

type TLSModule struct {
	config *tls.Config
}

func NoOpWrapper() func(net.Conn) net.Conn {
	return func(c net.Conn) net.Conn {
		return c
	}
}

func NewServer(addr string) (*Server, error) {
	res := Server{
		addr:     addr,
		listener: nil,
		wrapConn: nil,
	}

	return &res, nil
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.listener = ln
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			logger.Log{
				Msg:   "Error closing listener",
				Level: logger.ERROR,
				Fields: map[string]any{
					"error": err.Error(),
				},
			}.Print()

			return
		}
	}(ln)

	logger.Log{
		Msg:   "Listening on " + s.addr,
		Level: logger.INFO,
	}.Print()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		if s.wrapConn != nil {
			conn = s.wrapConn(conn)
			if conn == nil {
				continue
			}
		}

		go handler.HandleConnection(conn)
	}
}

func NewTLSModule(certFile, keyFile string) (*TLSModule, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return &TLSModule{
		config: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}, nil
}
func (t *TLSModule) Wrapper() func(net.Conn) net.Conn {
	return func(c net.Conn) net.Conn {
		_ = c.SetDeadline(time.Now().Add(10 * time.Second))

		tlsConn := tls.Server(c, t.config)

		err := tlsConn.Handshake()
		if err != nil {
			err := c.Close()
			if err != nil {
				return nil
			}
			return nil
		}

		_ = tlsConn.SetDeadline(time.Time{})

		return tlsConn
	}
}

func (s *Server) UseTLS(module *TLSModule) {
	s.wrapConn = module.Wrapper()
}
