package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const PORT = 8000

func main() {
	// Components routing:
	app.Route("/hello", func() app.Composer {
		return &hello{}
	})

	if app.IsClient {
		rpcUrl := fmt.Sprintf("http://%s/rpc", app.Window().URL().Host)
		_, err := jsonrpc.NewMergeClient(context.Background(), rpcUrl, "SimpleServerHandler", []any{&api}, nil, jsonrpc.WithHTTPClient(&http.Client{}))
		log.Println("jsonrpc.NewClient", err)
	}
	app.RunWhenOnBrowser()

	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("SimpleServerHandler", &API{})
	http.Handle("/rpc", enableCors(rpcServer))

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello RPC",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
