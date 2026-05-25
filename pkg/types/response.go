package types

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/server/status"
)

type Response struct {
	Type    ResType       `json:"type"`
	Status  status.Status `json:"status"`
	Info    Info          `json:"info"`
	Payload Payload       `json:"payload,omitempty"`
}

func (res *Response) ToJSON() string {
	marshal, _ := json.Marshal(res)
	return string(marshal)
}
