docker-build:
	docker build -t project-auction .

docker-run:
	docker run -p $(port) project-auction

docker-build-run: docker-build
	docker run -p $(port) project-auction

swag-fmt:
	swag fmt

swag-init: swag-fmt
	swag init -g cmd/main.go

goose-path:
	export GOOSE_MIGRATION_DIR=internal/adapters/postgres/migrations

goose-create:
	goose create $(name)_table sql

goose-up:
	goose postgres "postgres://test:test@localhost:5432/test" up

goose-down:
	goose postgres "postgres://test:test@localhost:5432/test" down

lint:
	golangci-lint run

run:
	go run cmd/main.go