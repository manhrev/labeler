build-proto:
	rm -rf ./pkg/api && cd api && buf generate && go mod tidy && go mod vendor

build-queries:
	rm -rf ./pkg/db && sqlc generate

run-local:
	export LISTEN_PORT=8080 && \
    export DB_USERNAME=postgres && \
    export DB_PASSWORD=password@1 && \
    export DB_HOST=localhost && \
    export DB_PORT=5432 && \
    export DB_NAME=backend_service && \
    export DB_SSL_MODE=disable && \
	export JWT_KEY=some_key_ok && \
	export ENVIRONMENT=LOCAL && \
	go run ./cmd/main.go
