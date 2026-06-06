package middleware

import (
	"context"

	"github.com/Sephy314/cnps/pkg/server/dto"
	"github.com/Sephy314/cnps/pkg/types"
)

var reqs = map[string]int{}

func RateLimit(n dto.Handler) dto.Handler {
	return func(c context.Context, req types.Request) (res types.Response, err error) {
		ip := c.Value("IP").(string)

		_, o := reqs[ip]
		if !o {
			reqs[ip] = 0
		}

		reqs[ip]++

		return n(c, req)
	}
}
