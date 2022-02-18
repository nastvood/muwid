package images

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-chi/render"

	"github.com/nastvood/muwid/internal/util"
)

type pullRequest struct {
	Repo string `json:"repo"`
}

func (p *pullRequest) Bind(r *http.Request) error {
	return nil
}

func list(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images, err := cli.ImageList(ctx, types.ImageListOptions{
			All: true,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		util.HttpJSON(w, images)
	}
}

func pull(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &pullRequest{}
		if err := render.Bind(r, data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reader, err := cli.ImagePull(ctx, data.Repo, types.ImagePullOptions{})
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer reader.Close()
		//nolint
		io.Copy(w, reader)
	}
}
