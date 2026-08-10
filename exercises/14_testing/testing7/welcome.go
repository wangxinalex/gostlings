// Package welcome demonstrates dependency injection through a small interface.
package welcome

import "context"

type UserStore interface {
	Name(ctx context.Context, id int) (string, error)
}

func Welcome(ctx context.Context, store UserStore, id int) (string, error) {
	name, err := store.Name(ctx, id)
	if err != nil {
		return "", err
	}
	return "welcome, " + name, nil
}
