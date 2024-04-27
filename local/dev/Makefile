PATH_TO_PROJECT=../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_FIREBASE_EMULATOR_SUITE_PORT=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev) \
	APP_PORT_APISVR_GRPC=$(APP_PORT_APISVR_GRPC_dev) \
	APP_PORT_UISVR_HTTP=$(APP_PORT_UISVR_HTTP_e2e_test)

.PHONY: envs
envs:
	@echo $(ENVS)

.PHONY: run
run:
	$(ENVS) docker compose up

.PHONY: up
up:
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down

DOCKER_COMPOSE_NAME=$(shell cat ./docker-compose.yml | yq .name)
DOCKER_COMPOSE_SERVICES=$(shell cat docker-compose.yml | yq '.services | select(.build.context != "") | keys | join(" ")')

.PHONY: rmi
rmi: down
	for service in $(DOCKER_COMPOSE_SERVICES); do \
		docker rmi $(DOCKER_COMPOSE_NAME)-$$service:latest; \
	done

.PHONY: rebuild
rebuild: rmi
	$(ENVS) docker compose build
