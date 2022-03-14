package movie

import (
	"context"
)

func (t *Movie) ClearCompleted(ctx context.Context) error {
	var movies []MovieItem
	if _, err := t.kv.FindAll(&movies); err != nil {
		return err
	}
	var toDelete []string
	for _, i := range movies {
		if i.Completed {
			toDelete = append(toDelete, i.ID)
		}
	}
	if err := t.kv.Delete(toDelete...); err != nil {
		return err
	}
	return nil
}
