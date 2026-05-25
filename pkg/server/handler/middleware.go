package handler

var Middlewares []Middleware

func Chain(final Handler, mws ...Middleware) Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		final = mws[i](final)
	}
	return final
}

func AddMiddlewares(mws ...Middleware) {
	Middlewares = append(Middlewares, mws...)
}
