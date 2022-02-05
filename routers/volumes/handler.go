package volumes

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"

	"github.com/nastvood/muwid/internals/util"
)

func list(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		volumes, err := cli.VolumeList(ctx, filters.NewArgs())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		util.HttpJSON(w, volumes)
	}
}

func remove(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		volumeID := chi.URLParam(r, "volumeID")
		err := cli.VolumeRemove(ctx, volumeID, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
