package response

import (
	"net"

	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/utils"
)

func WriteResponse(conn *net.Conn, res types.Response) {
	serialised := utils.SerialiseResponse(res)

	_, _ = (*conn).Write(append(serialised, byte('\n')))
}
