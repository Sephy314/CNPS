package errors

import (
	"github.com/Sephy314/cnps/pkg/server/status"
)

type CNPSError struct {
	Code    status.Status
	Message string
}

func (e *CNPSError) Error() string {
	return e.Message
}
