package main

import (
	"context"
	"fmt"
	"goweb/backend"
	"goweb/frontend"
	"log"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	// Components routing:
	app.Route("/hello", func() app.Composer {
		return &frontend.Hello{}
	})

	if app.IsClient {
		rpcUrl := fmt.Sprintf("http://%s/rpc", app.Window().URL().Host)
		_, err := jsonrpc.NewMergeClient(context.Background(), rpcUrl, "SimpleServerHandler", []any{
			&backend.APIClient,
		}, nil, jsonrpc.WithHTTPClient(&http.Client{}))
		if err != nil {
			log.Fatal(err)
		}
	}
	app.RunWhenOnBrowser()

	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("SimpleServerHandler", &backend.API{})
	http.Handle("/rpc", enableCors(rpcServer))

	http.Handle("/", &app.Handler{
		Name:        "Hello RPC",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
