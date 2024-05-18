docker-build:
	docker build -t project-auction .

docker-run:
	docker run -p $(port) project-auction

docker-build-run: docker-build
	docker run -p $(port) project-auction

swag-init:
	swag init -g cmd/main.go

swag-fmt:
	swag fmt

goose-path:
	export GOOSE_MIGRATION_DIR=migrations

goose-create:
	goose create $(name)_table sql

goose-up:
	goose postgres "postgres://test:test@localhost:5432/test" up

goose-down:
	goose postgres "postgres://test:test@localhost:5432/test" down

run:
	go run cmd/main.go