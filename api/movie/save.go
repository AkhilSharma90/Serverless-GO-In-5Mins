package movie

import (
	"context"
)

type SaveRequest struct {
	Movie MovieItem
}

func (t *Movie) Save(ctx context.Context, req *SaveRequest) error {
	return t.kv.Put(req.Movie.ID, req.Movie)
}
