package utils

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/types"
)

func SerialiseResponse(res types.Response) []byte {
	marshal, err := json.Marshal(res)
	if err != nil {
		return nil
	}
	return marshal
}
