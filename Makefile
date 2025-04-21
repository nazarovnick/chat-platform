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


