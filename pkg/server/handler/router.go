package handler

import (
	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/utils"
)

func Route(req string) (func(types.Request) (types.Response, error), error) {
	parsedReq, err := utils.ParseRequest(req)

	if err != nil {
		return nil, errors.BadRequestError("Error to parse request")
	}

	cmd := parsedReq.Cmd

	handler, ok := ROUTES.Routes[cmd]

	if !ok {
		return nil, errors.NotFoundError("Command Not Found")
	}

	return handler, nil

}
