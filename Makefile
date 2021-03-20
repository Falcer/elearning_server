run:
	@go run main.go
build:
	@go build -o ./build/main main.go
build-run:
	@go build -o ./build/main main.go && ./build/main
test:
	@go test ./...
