DOCKER_RUN=docker-compose -f ../docker/docker-compose.yml run --rm app

build:
	docker-compose -f ../docker/docker-compose.yml build
up:
	docker-compose -f ../docker/docker-compose.yml up -d
down:
	docker-compose -f ../docker/docker-compose.yml down
tidy:
	$(DOCKER_RUN) go mod tidy
wire-api:
	$(DOCKER_RUN) sh -c "cd src/cmd/api/ && wire"