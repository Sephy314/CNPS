package main

import (
	"fmt"
	"log"

	"github.com/Sephy314/cnps/pkg/dto"
	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/middleware"
	"github.com/Sephy314/cnps/pkg/server/router"
	"github.com/Sephy314/cnps/pkg/server/status"
	cnps "github.com/Sephy314/cnps/pkg/server/svr"
)

func main() {
	router.AddRoutes(".test", handler)

	middleware.AddMiddlewares(
		testMiddleware,
		testAnotherMiddleware,
	)

	err := cnps.Start(":31415")

	if err != nil {
		log.Fatal(err)
	}

}

func handler(req dto.Request) (dto.Response, error) {
	logger.Log{
		Msg:   fmt.Sprintf("Request: %+v", req),
		Level: logger.INFO,
	}.Print()

	return dto.Response{
		Type:    "RES",
		Status:  status.StatusOK,
		Payload: nil,
	}, nil
}

func testMiddleware(n router.Handler) router.Handler {
	return func(req dto.Request) (dto.Response, error) {
		logger.Log{
			Msg:   "Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(req)
		if err != nil {
			return dto.Response{}, err
		}

		return r, nil
	}
}

func testAnotherMiddleware(n router.Handler) router.Handler {
	return func(req dto.Request) (dto.Response, error) {
		logger.Log{
			Msg:   "Another Middleware Worked",
			Level: logger.DEBUG,
		}.Print()

		r, err := n(req)
		if err != nil {
			return dto.Response{}, err
		}

		return r, nil
	}
}
