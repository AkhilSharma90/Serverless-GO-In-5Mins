package movie

import (
	"context"
)

type GetResponse struct {
	Movies []MovieItem
}

func (t *Movie) Get(ctx context.Context) (*GetResponse, error) {
	var movies []MovieItem
	_, err := t.kv.FindAll(&movies)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		Movies: movies,
	}, nil
}
