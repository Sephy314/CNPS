package handler

import (
	"context"

	"github.com/Sephy314/cnps/pkg/types"
)

var ROUTES = NewRouters()

func NewRouters() Routers {
	return Routers{
		Routes: make(map[string]Handler),
	}
}

func AddRoutes(cmd string, handler Handler) {
	ROUTES.Routes[cmd] = handler
}

type Routers struct {
	Routes map[string]Handler
}

type Middleware func(next Handler) Handler
type Handler func(ctx context.Context, req types.Request) (types.Response, error)
