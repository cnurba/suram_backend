build:
	@go build -o bin/suram_backend  cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/suram_backend