SERVICE_NAME=auth
DOCKER_DIR=./docker

build-simple:
	docker build . -f $(DOCKER_DIR)/simple.Dockerfile -t $(SERVICE_NAME):simple

build-multi:
	docker build . -f $(DOCKER_DIR)/multi.Dockerfile -t $(SERVICE_NAME):multi

update-swagger-docs:
	swag init -g cmd/server/main.go -o docs/swagger
	sleep 1
	mv docs/swagger/swagger.yaml api/openapi/v1/swagger.yaml
	mv docs/swagger/swagger.json api/openapi/v1/swagger.json
