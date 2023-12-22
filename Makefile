build:
	@go build -o bin/go-crud

run: build
	@./bin/go-crud

test:
	@go test -v ./...