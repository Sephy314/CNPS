package handler

import (
	"log"

	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/utils"
)

func HandleRequest(msg string) (*types.Response, error) {
	handler, err := Route(msg)

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

	handler = Chain(handler, Middlewares...)

	res, err := handler(*parsedReq)

	if err != nil {
		log.Printf("Error handling request: %v", err)
		return nil, err
	}

	return &res, nil
}
