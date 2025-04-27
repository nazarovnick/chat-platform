SERVICE_NAME=auth
DOCKER_DIR=./docker

build-simple:
	docker build . -f $(DOCKER_DIR)/simple.Dockerfile -t $(SERVICE_NAME):simple

build-multi:
	docker build . -f $(DOCKER_DIR)/multi.Dockerfile -t $(SERVICE_NAME):multi