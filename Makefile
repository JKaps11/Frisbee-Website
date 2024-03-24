build:
	@go build -o bin/goFrisbee

run: build
	@./bin/goFrisbee

test:
	@go test -v ./...
