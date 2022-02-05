package images

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, cli *client.Client) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", list(ctx, cli))
		r.Post("/pull", pull(ctx, cli))
	}
}
