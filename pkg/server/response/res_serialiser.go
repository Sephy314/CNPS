package response

import (
	"encoding/json"

	"github.com/Sephy314/cnps/pkg/dto"
)

func SerialiseResponse(res dto.Response) []byte {
	marshal, err := json.Marshal(res)
	if err != nil {
		return nil
	}
	return marshal
}
