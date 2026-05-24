package router

import (
	dto2 "github.com/Sephy314/cnps/pkg/dto"
)

var ROUTES = NewRouters()

func NewRouters() Routers {
	return Routers{
		Routes: make(map[string]func(req dto2.Request) (dto2.Response, error)),
	}
}

func AddRoutes(cmd string, handler func(req dto2.Request) (dto2.Response, error)) {
	ROUTES.Routes[cmd] = handler
}

type Routers struct {
	Routes map[string]func(req dto2.Request) (dto2.Response, error)
}
