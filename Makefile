build:
	go build -o bin/syncstream-go-websocket ./cmd/server

run:
	go run ./cmd/server/main.go

test:
	go test -v ./...
