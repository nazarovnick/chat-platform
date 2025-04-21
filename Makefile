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