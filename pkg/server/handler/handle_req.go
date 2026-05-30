package handler

import (
	"context"
	"log"

	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/server/middleware"
	"github.com/Sephy314/cnps/pkg/server/route"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/utils"
)

func HandleRequest(c context.Context, msg string) (*types.Response, error) {
	handler, err := route.Route(msg)

	if err != nil {
		return nil, err
	}

	if handler == nil {
		log.Printf("Command Not Found: %v", msg)
		return nil, errors.NotFoundError("Command Not Found.")
	}

	parsedReq, err := utils.ParseRequest(msg)

	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return nil, err
	}

	handler = middleware.Chain(handler, middleware.Middlewares...)

	res, err := handler(c, *parsedReq)

	if err != nil {
		log.Printf("Error handling request: %v", err)
		return nil, err
	}

	return &res, nil
}
