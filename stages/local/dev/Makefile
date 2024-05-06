PATH_TO_PROJECT=../../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

APP_PORT_RPROXY_HTTP=10000

ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_FIREBASE_EMULATOR_SUITE_PORT=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev) \
	APP_PORT_APISVR=$(APP_PORT_APISVR_dev) \
	APP_PORT_UISVR=$(APP_PORT_UISVR_e2e_test) \
	APP_PORT_RPROXY_HTTP=$(APP_PORT_RPROXY_HTTP)

.PHONY: envs
envs:
	@echo $(ENVS)

DOCKER_COMPOSE_NAME=$(shell cat ./docker-compose.yml | yq .name)
DOCKER_COMPOSE_SERVICES=$(shell cat docker-compose.yml | yq '.services | select(.build.context != "") | keys | join(" ")')

DOCKER_COMPOSE_SERVICES_FOR_MIDDLEWARES=firebase-emulators
DOCKER_COMPOSE_SERVICES_FOR_APISVR=$(DOCKER_COMPOSE_SERVICES_FOR_MIDDLEWARES) apisvr
DOCKER_COMPOSE_SERVICES_FOR_UISVR=$(DOCKER_COMPOSE_SERVICES_FOR_APISVR) uisvr
DOCKER_COMPOSE_SERVICES_FOR_RPROXY=$(DOCKER_COMPOSE_SERVICES_FOR_UISVR) rproxy

.PHONY: validate_services
validate_services:
	@if [ "$(DOCKER_COMPOSE_SERVICES)" = "$(DOCKER_COMPOSE_SERVICES_FOR_RPROXY)" ] ;\
	then \
		echo "services are valid"; \
	else \
		echo "services are invalid"; \
		exit 1; \
	fi

.PHONY: setup
setup: validate_services
	$(MAKE) -C $(PATH_TO_UISVR) setup

.PHONY: run
run: setup
	$(ENVS) docker compose up

.PHONY: up
up: setup
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down


.PHONY: rmi
rmi: down
	for service in $(DOCKER_COMPOSE_SERVICES); do \
		docker rmi $(DOCKER_COMPOSE_NAME)-$$service:latest; \
	done

.PHONY: rebuild
rebuild:
	$(ENVS) $(MAKE) rmi ; \
	$(ENVS) docker compose build
