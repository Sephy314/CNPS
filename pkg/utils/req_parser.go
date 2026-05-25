package utils

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/types"
)

func ParseRequest(rawreq string) (*types.Request, error) {
	var req types.Request

	if err := json.Unmarshal([]byte(rawreq), &req); err != nil {
		return nil, err
	}

	if req.Type == "" ||
		req.Cmd == "" ||
		req.Act == "" ||
		req.Target == "" {
		return nil, errors.BadRequestError("Request Type is required")
	}

	return &req, nil
}
