run:
	@cd app
	@go run main.go
build:
	@cd app
	@go build -o ./build/main main.go
build-run:
	@cd app
	@go build -o ./build/main main.go && ./build/main
test:
	@cd app
	@go test ./...

reload-nginx:
	@docker-compose exec nginx nginx -s reload
