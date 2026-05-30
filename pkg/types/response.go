package types

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/types/status"
)

type Response struct {
	Status  status.Status `json:"status"`
	Info    Info          `json:"info"`
	Payload any           `json:"payload,omitempty"`
}

func (res *Response) ToJSON() string {
	marshal, _ := json.Marshal(res)
	return string(marshal)
}
