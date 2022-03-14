package movie

import (
	"context"
)

type DestroyRequest struct {
	ID string
}

func (t *Movie) Destroy(ctx context.Context, req *DestroyRequest) error {
	return t.kv.Delete(req.ID)
}
