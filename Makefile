deploy:
	go run github.com/99designs/gqlgen generate
	go mod tidy
	go run ./cmd/server/server.go
