package middleware

import (
	"context"

	"github.com/Sephy314/cnps/pkg/logger"
	"github.com/Sephy314/cnps/pkg/server/dto"
	"github.com/Sephy314/cnps/pkg/types"
)

func Recovery(next dto.Handler) dto.Handler {
	return func(c context.Context, req types.Request) (res types.Response, err error) {

		defer func() {
			if r := recover(); r != nil {

				logger.Log{
					Msg:   "CNPS panic recovered",
					Level: logger.ERROR,
					Fields: map[string]any{
						"cmd":   req.Cmd,
						"act":   req.Act,
						"panic": r,
					},
				}.Print()

				res = types.Response{
					Status: 40,
					Payload: map[string]any{
						"error": "internal server error",
					},
				}

				err = nil
			}
		}()

		return next(c, req)
	}
}
