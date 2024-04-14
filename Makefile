docker-build:
	docker build -t project-auction .

docker-run:
	docker run -p $(port) project-auction

docker-build-run: docker-build
	docker run -p $(port) project-auction

swag-init:
	swag init -g cmd/main.go

run:
	go run cmd/main.go