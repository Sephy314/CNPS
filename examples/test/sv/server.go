package main

import (
	"fmt"

	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/handler"
	cnps "github.com/Sephy314/cnps/pkg/server/svr"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/types/status"
)

func main() {
	handler.AddRoutes(".test", testHandler)

	handler.AddMiddlewares(
		testMiddleware,
		testAnotherMiddleware,
	)

	err := cnps.Start(":31415")

	if err != nil {
		logger.Log{
			Msg:   err,
			Level: logger.ERROR,
		}.Print()
	}

}

func testHandler(req types.Request) (types.Response, error) {
	logger.Log{
		Msg:   fmt.Sprintf("Handler got : %+v", req),
		Level: logger.INFO,
	}.Print()

	return types.Response{
		Status:  status.StatusOK,
		Payload: nil,
	}, nil
}

func testMiddleware(n handler.Handler) handler.Handler {
	return func(req types.Request) (types.Response, error) {
		logger.Log{
			Msg:   "Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(req)
		if err != nil {
			return types.Response{}, err
		}

		return r, nil
	}
}

func testAnotherMiddleware(n handler.Handler) handler.Handler {
	return func(req types.Request) (types.Response, error) {
		logger.Log{
			Msg:   "Another Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(req)
		if err != nil {
			return types.Response{}, err
		}

		return r, nil
	}
}
