# Start all containers in detached mode
up:
	docker-compose up -d

# Build and start all containers (forces rebuild)
up-build:
	docker-compose up -d --build

# Stop and remove all containers
down:
	docker-compose down -d

# Remove containers and ALL images (including pulled ones)
clean-images-all:
	docker-compose down --rmi all

# Remove containers and locally built images (from "build:" blocks)
clean-images-local:
	docker-compose down --rmi local


# Show live logs from all services
logs:
	docker-compose logs -f

# Restart all containers (hard reset)
restart:
	docker-compose down && docker-compose up -d

# Stop containers without removing them
stop:
	docker-compose stop


docker-env:
	@echo "Detecting OS and setting Docker env for Minikube..."
	@if [ "$$(uname -s)" = "Linux" ] || [ "$$(uname -s)" = "Darwin" ]; then \
		echo "Running on Linux/macOS"; \
		eval $$(minikube docker-env); \
	else \
		echo "You're on Windows. Please run this manually in PowerShell:"; \
		echo "  minikube docker-env | Invoke-Expression"; \
	fi


# =======================================
# AUTH SERVICE K8S COMMANDS
# =======================================

# Build the auth:blue image directly into Minikube's Docker
build-auth-blue:
	docker build -t auth:blue -f ./services/auth/docker/multi.Dockerfile ./services/auth

# Build the auth:green image directly into Minikube's Docker
build-auth-green:
	docker build -t auth:green -f ./services/auth/docker/multi.Dockerfile ./services/auth

# Deploy blue and green versions of the auth service along with its Kubernetes service
auth-deploy:
	kubectl apply -f k8s/auth/blue-deployment.yaml
	kubectl apply -f k8s/auth/green-deployment.yaml
	kubectl apply -f k8s/auth/service.yaml

# Switch traffic to the blue version
auth-switch-blue:
	kubectl patch service auth-service -p '{"spec": {"selector": {"app": "auth", "version": "blue"}}}'

# Switch traffic to the green version
auth-switch-green:
	kubectl patch service auth-service -p '{"spec": {"selector": {"app": "auth", "version": "green"}}}'


# =======================================
# CHAT SERVICE K8S COMMANDS
# =======================================

# Build the chat:blue image directly into Minikube's Docker
build-chat-blue:
	docker build -t chat:blue -f ./services/chat/docker/multi.Dockerfile ./services/chat

# Build the chat:green image directly into Minikube's Docker
build-chat-green:
	docker build -t chat:green -f ./services/chat/docker/multi.Dockerfile ./services/chat

# Deploy blue and green versions of the chat service along with its Kubernetes service
chat-deploy:
	kubectl apply -f k8s/chat/blue-deployment.yaml
	kubectl apply -f k8s/chat/green-deployment.yaml
	kubectl apply -f k8s/chat/service.yaml

# Switch traffic to the blue version
chat-switch-blue:
	kubectl patch service chat-service -p '{"spec": {"selector": {"app": "chat", "version": "blue"}}}'

# Switch traffic to the green version
chat-switch-green:
	kubectl patch service chat-service -p '{"spec": {"selector": {"app": "chat", "version": "green"}}}'


# =======================================
# FRIENDSHIP SERVICE K8S COMMANDS
# =======================================

# Build the friendship:blue image directly into Minikube's Docker
build-friendship-blue:
	docker build -t friendship:blue -f ./services/friendship/docker/multi.Dockerfile ./services/friendship

# Build the friendship:green image directly into Minikube's Docker
build-friendship-green:
	docker build -t friendship:green -f ./services/friendship/docker/multi.Dockerfile ./services/friendship

# Deploy blue and green versions of the friendship service along with its Kubernetes service
friendship-deploy:
	kubectl apply -f k8s/friendship/blue-deployment.yaml
	kubectl apply -f k8s/friendship/green-deployment.yaml
	kubectl apply -f k8s/friendship/service.yaml

# Switch traffic to the blue version
friendship-switch-blue:
	kubectl patch service friendship-service -p '{"spec": {"selector": {"app": "friendship", "version": "blue"}}}'

# Switch traffic to the green version
friendship-switch-green:
	kubectl patch service friendship-service -p '{"spec": {"selector": {"app": "friendship", "version": "green"}}}'


# =======================================
# USER SERVICE K8S COMMANDS
# =======================================

# Build the user:blue image directly into Minikube's Docker
build-user-blue:
	docker build -t user:blue -f ./services/user/docker/multi.Dockerfile ./services/user

# Build the user:green image directly into Minikube's Docker
build-user-green:
	docker build -t user:green -f ./services/user/docker/multi.Dockerfile ./services/user

# Deploy blue and green versions of the user service along with its Kubernetes service
user-deploy:
	kubectl apply -f k8s/user/blue-deployment.yaml
	kubectl apply -f k8s/user/green-deployment.yaml
	kubectl apply -f k8s/user/service.yaml

# Switch traffic to the blue version
user-switch-blue:
	kubectl patch service user-service -p '{"spec": {"selector": {"app": "user", "version": "blue"}}}'

# Switch traffic to the green version
user-switch-green:
	kubectl patch service user-service -p '{"spec": {"selector": {"app": "user", "version": "green"}}}'
