#! /bin/bash
GOARCH=wasm GOOS=js go build -o web/app.wasm
(sleep 2 && open 'http://localhost:8000') &
go run .