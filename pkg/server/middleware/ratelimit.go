package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/Sephy314/cnps/pkg/server/dto"
	"github.com/Sephy314/cnps/pkg/types"
	"github.com/Sephy314/cnps/pkg/types/status"
)

func RateLimit(rpm int) func(n dto.Handler) dto.Handler {

	var reqs = map[string]int{}
	var reqMtx sync.Mutex

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			reqMtx.Lock()
			reqs = make(map[string]int)
			reqMtx.Unlock()
		}
	}()

	return func(n dto.Handler) dto.Handler {

		return func(c context.Context, req types.Request) (res types.Response, err error) {
			ip := c.Value("IP").(string)

			reqMtx.Lock()

			_, o := reqs[ip]
			if !o {
				reqs[ip] = 0
			}

			reqs[ip]++

			reqMtx.Unlock()

			if reqs[ip] >= rpm {
				return types.Response{
					Status:  status.StatusTooManyRequests,
					Info:    types.Info{},
					Payload: nil,
				}, nil
			}

			return n(c, req)

		}
	}
}
