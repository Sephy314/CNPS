package middleware

import (
	"github.com/Sephy314/cnps/pkg/server/router"
)

var Middlewares []router.Middleware

func Chain(final router.Handler, mws ...router.Middleware) router.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		final = mws[i](final)
	}
	return final
}

func AddMiddlewares(mws ...router.Middleware) {
	Middlewares = append(Middlewares, mws...)
}
