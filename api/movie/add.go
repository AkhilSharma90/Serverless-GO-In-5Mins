package movie

import (
	"context"

	"github.com/google/uuid"
)

type AddRequest struct {
	Title string
}

func (t *Movie) Add(ctx context.Context, req *AddRequest) error {
	id := uuid.NewString()
	return t.kv.Put(id, &MovieItem{
		ID:        id,
		Title:     req.Title,
		Completed: false,
	})
}
