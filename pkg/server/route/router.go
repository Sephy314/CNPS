package route

import (
	"github.com/Sephy314/cnps/pkg/server/dto"
)

var ROUTES = NewRouters()

func NewRouters() Routers {
	return Routers{
		Routes: make(map[string]dto.Handler),
	}
}

func AddRoutes(cmd string, handler dto.Handler) {
	ROUTES.Routes[cmd] = handler
}

type Routers struct {
	Routes map[string]dto.Handler
}
