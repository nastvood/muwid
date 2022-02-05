package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/nastvood/muwid/routers/containers"
	"github.com/nastvood/muwid/routers/images"
	"github.com/nastvood/muwid/routers/volumes"
)

const defaultAddr = ":3001"

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalln(err)
	}

	pAddress := flag.String("addr", defaultAddr, fmt.Sprintf("TCP network address (default %s)", defaultAddr))
	flag.Parse()
	address := defaultAddr
	if pAddress != nil {
		address = *pAddress
	}

	log.Println("Run server", address)

	err = http.ListenAndServe(address, chiInit(ctx, cli))
	if err != nil {
		log.Fatalln(err)
	}
}

func chiInit(ctx context.Context, cli *client.Client) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/containers", containers.Router(ctx, cli))
	r.Route("/images", images.Router(ctx, cli))
	r.Route("/volumes", volumes.Router(ctx, cli))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//nolint
		w.Write(data)
	})

	return r
}

//docker system df
