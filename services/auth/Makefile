SERVICE_NAME=auth
DOCKER_DIR=$(CURDIR)docker
LOCAL_BIN=$(CURDIR)/bin
SWAG= PATH="$$PATH:$(LOCAL_BIN)" swag


build-simple:
	docker build . -f $(DOCKER_DIR)/simple.Dockerfile -t $(SERVICE_NAME):simple



build-multi:
	docker build . -f $(DOCKER_DIR)/multi.Dockerfile -t $(SERVICE_NAME):multi

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/swaggo/swag/cmd/swag@latest

.generate-swagger-docs:
	$(SWAG) init -g ./cmd/server/main.go -o docs/swagger
	sleep 1
	mv docs/swagger/swagger.yaml api/openapi/v1/swagger.yaml
	mv docs/swagger/swagger.json api/openapi/v1/swagger.json

update-swagger: .bin-deps .generate-swagger-docs