package main

import (
	"context"
	"fmt"
	"goweb/backend/models"
	"goweb/backend/services"
	"goweb/frontend"
	"log"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/klauspost/compress/gzhttp"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Register your new client here
	clients := map[string]any{
		"Notes": &services.NotesClient,
		"Clock": &services.ClockClient,
	}

	frontend.RegisterRoutes()
	if app.IsClient {
		rpcUrl := fmt.Sprintf("http://%s/rpc", app.Window().URL().Host)
		for name, client := range clients {
			_, err := jsonrpc.NewMergeClient(context.Background(), rpcUrl, name, []any{
				client,
			}, nil, jsonrpc.WithHTTPClient(&http.Client{}))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	app.RunWhenOnBrowser()
	if !app.IsServer {
		return
	}

	db := initializeDatabase()
	// 2. Register your new service here
	services := map[string]any{
		"Notes": services.NewNotesService(db),
		"Clock": services.NewClockService(),
	}
	initializeRPC(services)

	http.Handle("/", gzhttp.GzipHandler(&app.Handler{
		Name:        "Hello RPC",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/dark.min.css",
		},
	}))

	log.Println("Server started on http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 3. Add an AutoMigrate if you create a new model
	db.AutoMigrate(&models.Note{})
	return db
}

func initializeRPC(services map[string]any) {
	rpcServer := jsonrpc.NewServer()
	for name, service := range services {
		rpcServer.Register(name, service)
	}
	http.Handle("/rpc", enableCors(rpcServer))
}
