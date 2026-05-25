package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Sephy314/cnps/pkg/server/error"
	"github.com/Sephy314/cnps/pkg/server/logger"
	response2 "github.com/Sephy314/cnps/pkg/server/response"
	"github.com/Sephy314/cnps/pkg/server/status"
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

			// CNPS error check
			var cnpErr *cnpserr.CNPSError
			if errors.As(err, &cnpErr) {
				log.Printf("CNPS Error Occurred: %v", cnpErr)

				newCnpsErr := response2.CreateCnpsErrorResponse(*cnpErr)

				cnpsLog := logger.ResponseLog{
					Log: logger.Log{
						Msg:   cnpErr.Message,
						Level: logger.ERROR,
					},
					ReqID:  reqId,
					Status: cnpErr.Code,
				}

				cnpsLog.Print()

				response2.WriteResponse(&conn, newCnpsErr)

				continue
			}

			// generic error
			log.Printf("Internal Error: %v", err)
			internalErr := response2.CreateErrorResponse(err)

			cnpsLog := logger.ResponseLog{
				Log: logger.Log{
					Msg:   err.Error(),
					Level: logger.ERROR,
				},
				ReqID:  reqId,
				Status: status.StatusInternalError,
			}

			cnpsLog.Print()

			response2.WriteResponse(&conn, internalErr)

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

			response2.WriteResponse(&conn, *res)
		}

	}
}
