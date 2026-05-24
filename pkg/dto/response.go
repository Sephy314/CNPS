package dto

import (
	"github.com/Sephy314/cnps/pkg/server/status"
)

type Response struct {
	Type    ResType       `json:"type"`
	Info    Info          `json:"info"`
	Status  status.Status `json:"status"`
	Payload Payload       `json:"payload,omitempty"`
}
