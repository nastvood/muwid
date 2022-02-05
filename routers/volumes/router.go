package volumes

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, cli *client.Client) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", list(ctx, cli))
		r.Get("/remove", remove(ctx, cli))
	}
}
