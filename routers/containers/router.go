package containers

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, cli *client.Client) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", list(ctx, cli))

		r.Get("/start/{containerID}", start(ctx, cli))
		r.Get("/stop/{containerID}", stop(ctx, cli))
		r.Get("/pause/{containerID}", pause(ctx, cli))
		r.Get("/unpause/{containerID}", unpause(ctx, cli))
		r.Get("/remove/{containerID}/{purge}", remove(ctx, cli))
	}
}
