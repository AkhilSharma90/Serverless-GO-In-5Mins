package movie

import (
	"context"
)

type ToggleRequest struct {
	ID  string
	Val bool
}

func (t *Movie) Toggle(ctx context.Context, req *ToggleRequest) error {
	i := &MovieItem{}
	if err := t.kv.Get(req.ID, i); err != nil {
		return err
	}
	i.Completed = req.Val
	return t.kv.Put(i.ID, i)
}
