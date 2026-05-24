package router

import (
	dto2 "github.com/Sephy314/cnps/pkg/dto"
	"github.com/Sephy314/cnps/pkg/server/error"
	"github.com/Sephy314/cnps/pkg/server/request"
)

func Route(req string) (func(dto2.Request) (dto2.Response, error), error) {
	parsedReq, err := request.ParseRequest(req)

	if err != nil {
		return nil, cnpserr.BadRequestError("Error to parse request")
	}

	cmd := parsedReq.Cmd

	handler, ok := ROUTES.Routes[cmd]

	if !ok {
		return nil, cnpserr.NotFoundError("Command Not Found")
	}

	return handler, nil

}
