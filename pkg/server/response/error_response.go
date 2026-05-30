package response

import (
	"github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/types/status"
)

func CreateCnpsErrorResponse(err errors.CNPSError) types.Response {
	res := types.Response{
		Status:  err.Code,
		Payload: nil,
	}

	return res
}

func CreateErrorResponse(err error) types.Response {
	res := types.Response{
		Status: status.StatusInternalError,
		Payload: map[string]interface{}{
			"error": err.Error(),
		},
	}
	return res
}
