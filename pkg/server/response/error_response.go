package response

import (
	dto2 "github.com/Sephy314/cnps/pkg/dto"
	cnperr "github.com/Sephy314/cnps/pkg/server/error"
	"github.com/Sephy314/cnps/pkg/server/status"
)

func CreateCnpsErrorResponse(err cnperr.CNPSError) dto2.Response {
	res := dto2.Response{
		Type:    dto2.ResTypeEr,
		Status:  err.Code,
		Payload: nil,
	}

	return res
}

func CreateErrorResponse(err error) dto2.Response {
	res := dto2.Response{
		Type:    dto2.ResTypeEr,
		Status:  status.StatusInternalError,
		Payload: nil,
	}
	return res
}
