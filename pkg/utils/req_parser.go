package utils

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/dto"
	"github.com/Sephy314/cnps/pkg/server/error"
)

func ParseRequest(rawreq string) (*dto.Request, error) {
	var res dto.Request

	if err := json.Unmarshal([]byte(rawreq), &res); err != nil {
		return nil, err
	}

	if res.Type == "" ||
		res.Cmd == "" ||
		res.Act == "" ||
		res.Target == "" {
		return nil, cnpserr.BadRequestError("Request Type is required")
	}

	return &res, nil
}
