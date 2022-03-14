package movie

import (
	"context"
)

type SaveRequest struct {
	Movie MovieItem
}

func (t *Movie) Update(ctx context.Context, req *SaveRequest) error {
	return t.kv.Put(req.Movie.ID, req.Movie)
}
