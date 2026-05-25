package main

import (
	"fmt"
	"log"

	"github.com/Sephy314/cnps/pkg/dto"
	"github.com/Sephy314/cnps/pkg/server/logger"
	"github.com/Sephy314/cnps/pkg/server/router"
	"github.com/Sephy314/cnps/pkg/server/status"
	cnps "github.com/Sephy314/cnps/pkg/server/svr"
)

func main() {
	router.AddRoutes(".test", handler)

	err := cnps.Start(":31415")

	if err != nil {
		log.Fatal(err)
	}

}

func handler(req dto.Request) (dto.Response, error) {
	logger.Log{
		Msg:   fmt.Sprintf("Request: %+v", req),
		Level: logger.INFO,
	}.Print()

	return dto.Response{
		Type:    "RES",
		Status:  status.StatusOK,
		Payload: nil,
	}, nil
}
