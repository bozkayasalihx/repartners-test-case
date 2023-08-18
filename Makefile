test: 
	@go clean -testcache && go test ./... -v

build: 
	@go build -o bin/calculatepacks ./internal/api

run: build
	@bin/calculatepacks
