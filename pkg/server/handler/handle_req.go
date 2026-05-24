package handler

import (
	"log"

	"github.com/Sephy314/cnps/pkg/dto"
	cnperr "github.com/Sephy314/cnps/pkg/server/error"
	"github.com/Sephy314/cnps/pkg/server/request"
	"github.com/Sephy314/cnps/pkg/server/router"
)

func HandleRequest(msg string) (*dto.Response, error) {
	handler, err := router.Route(msg)

	if err != nil {
		return nil, err
	}

	if handler == nil {
		log.Printf("Command Not Found: %v", msg)
		return nil, cnperr.NotFoundError("Command Not Found.")
	}

	parsedReq, err := request.ParseRequest(msg)

	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return nil, err
	}

	res, err := handler(*parsedReq)

	if err != nil {
		log.Printf("Error handling request: %v", err)
		return nil, err
	}

	return &res, nil
}
