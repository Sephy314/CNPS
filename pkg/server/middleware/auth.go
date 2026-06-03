package middleware

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/Sephy314/cnps/pkg/server/dto"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/types/status"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(sec []byte, p []string, a jwt.SigningMethod) func(dto.Handler) dto.Handler {
	return func(n dto.Handler) dto.Handler {
		return func(ctx context.Context, req types.Request) (types.Response, error) {
			if slices.Contains(p, req.Cmd) {
				return n(ctx, req)
			}

			hd := *req.Info.Ext

			tk := hd["authorization"].(string)

			if !strings.HasPrefix(tk, "Bearer ") {
				return types.Response{
					Status:  status.StatusUnauthorized,
					Info:    types.Info{},
					Payload: nil,
				}, nil
			}

			tk = strings.TrimPrefix(tk, "Bearer ")

			ps, e := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
				if token.Method != a {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return sec, nil
			})

			if e != nil || !ps.Valid {
				return types.Response{
					Status:  status.StatusUnauthorized,
					Info:    types.Info{},
					Payload: nil,
				}, nil
			}

			claims, ok := ps.Claims.(jwt.MapClaims)
			if !ok {
				return types.Response{
					Status:  status.StatusUnauthorized,
					Info:    types.Info{},
					Payload: nil,
				}, nil
			}

			sbj, err := claims.GetSubject()

			if err != nil {
				return types.Response{
					Status:  status.StatusUnauthorized,
					Info:    types.Info{},
					Payload: nil,
				}, nil
			}

			ctx = context.WithValue(
				ctx,
				"uid",
				sbj,
			)

			res, err := n(ctx, req)

			return res, err
		}
	}
}
