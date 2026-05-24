package response

import (
	"net"

	"github.com/Sephy314/cnps/pkg/dto"
)

func WriteResponse(conn *net.Conn, res dto.Response) {
	serialised := SerialiseResponse(res)

	_, _ = (*conn).Write(append(serialised, byte('\n')))
}
