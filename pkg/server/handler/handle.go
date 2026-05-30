package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Sephy314/cnps/pkg/logger"
	cnpserr "github.com/Sephy314/cnps/pkg/server/errors"
	"github.com/Sephy314/cnps/pkg/server/response"
	"github.com/Sephy314/cnps/pkg/types/status"
	"github.com/google/uuid"
)

func HandleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//logger.Log{
				//	Msg:   "Connection closed by remote hose",
				//	Level: logger.INFO,
				//}.Print()
			} else {
				logger.Log{
					Msg:   fmt.Sprintf("Error reading from connection: %v", err),
					Level: logger.ERROR,
				}.Print()
			}
			return
		}

		if msg == "\n" || msg == "" {
			continue
		}

		// Requested
		reqId := uuid.New().String()

		res, err := HandleRequest(msg)

		if err != nil {

			// CNPS errs check
			if cnpErr, ok := errors.AsType[*cnpserr.CNPSError](err); ok {
				log.Printf("CNPS Error Occurred: %v", cnpErr)

				newCnpsErr := response.CreateCnpsErrorResponse(*cnpErr)

				cnpsLog := logger.ResponseLog{
					Log: logger.Log{
						Msg:   cnpErr.Message,
						Level: logger.ERROR,
					},
					ReqID:  reqId,
					Status: cnpErr.Code,
				}

				cnpsLog.Print()

				response.WriteResponse(&conn, newCnpsErr)

				continue
			}

			// generic errs
			log.Printf("Internal Error: %v", err)
			internalErr := response.CreateErrorResponse(err)

			cnpsLog := logger.ResponseLog{
				Log: logger.Log{
					Msg:   err.Error(),
					Level: logger.ERROR,
				},
				ReqID:  reqId,
				Status: status.StatusInternalError,
			}

			cnpsLog.Print()

			response.WriteResponse(&conn, internalErr)

			return
		}

		if res != nil {
			cnpsLog := logger.ResponseLog{
				Log: logger.Log{
					Msg:   "Requested",
					Level: logger.INFO,
				},
				ReqID:  reqId,
				Status: res.Status,
			}

			cnpsLog.Print()

			response.WriteResponse(&conn, *res)
		}

	}
}
