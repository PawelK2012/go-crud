build:
	@go build -o bin/go-crud

run: build
	@./bin/go-crud

test:
	@go test -v ./...

cover:
	go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out