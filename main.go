package main

import (
	"context"
	"fmt"
	"goweb/backend"
	"goweb/frontend"
	"log"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/klauspost/compress/gzhttp"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

var SERVICES = map[string]Service{
	"Notes": {
		Service: &backend.Notes{},
		Client:  &backend.NotesClient,
	},
	"Clock": {
		Service: &backend.Clock{},
		Client:  &backend.ClockClient,
	},
}

type Service struct {
	Service interface{}
	Client  interface{}
}

func main() {
	frontend.RegisterRoutes()

	if app.IsClient {
		rpcUrl := fmt.Sprintf("http://%s/rpc", app.Window().URL().Host)
		for name, service := range SERVICES {
			_, err := jsonrpc.NewMergeClient(context.Background(), rpcUrl, name, []any{
				service.Client,
			}, nil, jsonrpc.WithHTTPClient(&http.Client{}))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	app.RunWhenOnBrowser()

	rpcServer := jsonrpc.NewServer()
	for name, service := range SERVICES {
		rpcServer.Register(name, service.Service)
	}
	http.Handle("/rpc", enableCors(rpcServer))

	http.Handle("/", gzhttp.GzipHandler(&app.Handler{
		Name:        "Hello RPC",
		Description: "An Hello World! example",
	}))

	log.Println("Server started on http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
