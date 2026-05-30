package errors

import (
	"github.com/Sephy314/cnps/pkg/types/status"
)

func NotFoundError(msg string) error {
	return &CNPSError{
		Code:    status.StatusNotFound,
		Message: msg,
	}
}

//func InternalError(msg string) error {
//	return &CNPSError{
//		Code:    status.StatusInternalError,
//		Message: msg,
//	}
//}

func BadRequestError(msg string) error {
	return &CNPSError{
		Code:    status.StatusBadRequest,
		Message: msg,
	}
}
