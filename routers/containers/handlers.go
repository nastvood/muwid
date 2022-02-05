package containers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"

	"github.com/nastvood/muwid/internals/util"
)

func list(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
			All: true,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		util.HttpJSON(w, containers)
	}
}

func start(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "containerID")
		err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func stop(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "containerID")
		timeout := time.Second * 5
		err := cli.ContainerStop(ctx, containerID, &timeout)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func pause(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "containerID")
		err := cli.ContainerPause(ctx, containerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func unpause(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "containerID")
		err := cli.ContainerUnpause(ctx, containerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func remove(ctx context.Context, cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "containerID")
		purgeS := chi.URLParam(r, "purge")

		opt := types.ContainerRemoveOptions{
			Force: true,
		}
		if purge, _ := strconv.ParseBool(purgeS); purge {
			opt.RemoveLinks = true
			opt.RemoveVolumes = true
		}

		err := cli.ContainerRemove(ctx, containerID, opt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
