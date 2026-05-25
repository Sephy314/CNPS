package response

import (
	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/server/status"
	"github.com/Sephy314/cnps/pkg/types"
)

func CreateCnpsErrorResponse(err errors.CNPSError) types.Response {
	res := types.Response{
		Type:    types.ResTypeEr,
		Status:  err.Code,
		Payload: nil,
	}

	return res
}

func CreateErrorResponse(err error) types.Response {
	res := types.Response{
		Type:   types.ResTypeEr,
		Status: status.StatusInternalError,
		Payload: map[string]interface{}{
			"error": err.Error(),
		},
	}
	return res
}
