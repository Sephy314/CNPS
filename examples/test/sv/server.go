package main

import (
	"context"
	"fmt"

	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/dto"
	"github.com/Sephy314/cnps/pkg/server/middleware"
	"github.com/Sephy314/cnps/pkg/server/route"
	cnps "github.com/Sephy314/cnps/pkg/server/svr"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/types/status"
)

func main() {
	route.AddRoutes(".test", testHandler)
	route.AddRoutes(".panic", testPanic)

	middleware.AddMiddlewares(
		testMiddleware,
		testAnotherMiddleware,
		middleware.Recovery,
	)

	err := cnps.Start(":31415")

	if err != nil {
		logger.Log{
			Msg:   err,
			Level: logger.ERROR,
		}.Print()
	}

}

func testHandler(_ context.Context, req types.Request) (types.Response, error) {
	logger.Log{
		Msg:   fmt.Sprintf("Handler got : %+v", req),
		Level: logger.INFO,
	}.Print()

	return types.Response{
		Status:  status.StatusOK,
		Payload: nil,
	}, nil
}

func testPanic(_ context.Context, _ types.Request) (types.Response, error) {
	panic("WOAHHHH")
}

func testMiddleware(n dto.Handler) dto.Handler {
	return func(c context.Context, req types.Request) (types.Response, error) {
		logger.Log{
			Msg:   "Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(c, req)
		if err != nil {
			return types.Response{}, err
		}

		return r, nil
	}
}

func testAnotherMiddleware(n dto.Handler) dto.Handler {
	return func(c context.Context, req types.Request) (types.Response, error) {
		logger.Log{
			Msg:   "Another Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(c, req)
		if err != nil {
			return types.Response{}, err
		}

		return r, nil
	}
}
