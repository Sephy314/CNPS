package dto

import (
	"context"

	"github.com/Sephy314/cnps/pkg/types"
)

type Handler func(ctx context.Context, req types.Request) (types.Response, error)
