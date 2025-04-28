# Name of the service (used as Docker image tag)
SERVICE_NAME=chat

# Path to the directory containing Dockerfiles
DOCKER_DIR=./docker

# Use the bin directory in the current working directory to install protoc plugins
LOCAL_BIN := $(CURDIR)/bin

# Prefix for the API version directory structure under ./api/proto/ (e.g., v1, v2, etc.)
# Used to build the full path to the .proto files for a specific version of the API.
API_VERSION := v1

# Add the local bin directory to PATH when running protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Path to the protobuf files
PROTO_PATH := $(CURDIR)/api/proto/$(API_VERSION)/$(SERVICE_NAME)

# Path where the generated .pb.go files will be placed
PKG_PROTO_PATH := $(CURDIR)/pkg/api/grpc

# Install the required protoc plugins
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate .go files from .proto files using protoc
.protoc-generate:
	mkdir -p $(PKG_PROTO_PATH)
	$(PROTOC) --proto_path=$(PROTO_PATH) \
	--go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

# Run go mod tidy to clean up go.mod and go.sum
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy


# Generate code from protobuf definitions
generate: .bin-deps .protoc-generate .tidy

# Build the application binaries
build:
	go build -o $(LOCAL_BIN) ./cmd/server

# Build a simple Docker image using simple.Dockerfile
build-simple:
	docker build . -f $(DOCKER_DIR)/simple.Dockerfile -t $(SERVICE_NAME):simple

# Build a multi-stage Docker image using multi.Dockerfile
build-multi:
	docker build . -f $(DOCKER_DIR)/multi.Dockerfile -t $(SERVICE_NAME):multi

# Declare these targets as phony to prevent Make from checking for corresponding files
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	generate \
	build