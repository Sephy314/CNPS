package middleware

import (
	"github.com/Sephy314/cnps/pkg/server/dto"
)

var Middlewares []Middleware

func Chain(final dto.Handler, mws ...Middleware) dto.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		final = mws[i](final)
	}
	return final
}

func AddMiddlewares(mws ...Middleware) {
	Middlewares = append(Middlewares, mws...)
}

type Middleware func(next dto.Handler) dto.Handler
