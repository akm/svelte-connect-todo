PATH_TO_ROOT=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/makefiles/template/text-template-cli.mk

DEV_TARGET?=all

BASE_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_PORT_FIREBASE_AUTH=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_PORT_FIREBASE_EMULATOR_SUITE=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev) \
	APP_PORT_APISVR=$(APP_PORT_APISVR_dev) \
	APP_PORT_UISVR=$(APP_PORT_UISVR_e2e_test) \
	APP_PORT_RPROXY=$(APP_PORT_RPROXY_dev) \
	DEV_TARGET=$(DEV_TARGET)

APP_UISVR_DOT_ENV_PATH=./uisvr.env
$(APP_UISVR_DOT_ENV_PATH): uisvr_env_gen
.PHONY: uisvr_env_gen
uisvr_env_gen: $(TEXT_TEMPLATE_CLI)
	$(BASE_ENVS) $(TEXT_TEMPLATE_CLI) ./uisvr.env.template > $(APP_UISVR_DOT_ENV_PATH)

APP_RPROXY_ENVOY_YAML=./envoy.yaml
$(APP_RPROXY_ENVOY_YAML): envoy_yaml_gen
.PHONY: envoy_yaml_gen
envoy_yaml_gen: $(TEXT_TEMPLATE_CLI)
	$(BASE_ENVS) $(TEXT_TEMPLATE_CLI) envoy.yaml.template > $(APP_RPROXY_ENVOY_YAML)

ENVS=\
	$(BASE_ENVS) \
	APP_UISVR_DOT_ENV=$(APP_UISVR_DOT_ENV_PATH) \
	APP_RPROXY_ENVOY_YAML=$(APP_RPROXY_ENVOY_YAML)

.PHONY: envs
envs:
	@echo $(ENVS)

DOCKER_COMPOSE_NAME=$(shell cat ./docker-compose.yml | yq .name)
DOCKER_COMPOSE_SERVICES=$(shell cat docker-compose.yml | yq '.services | select(.build.context != "") | keys | join(" ")')

DOCKER_COMPOSE_SERVICES_FOR_DEV_apisvr=firebase-emulators
DOCKER_COMPOSE_SERVICES_FOR_DEV_uisvr=$(DOCKER_COMPOSE_SERVICES_FOR_DEV_apisvr) apisvr
DOCKER_COMPOSE_SERVICES_FOR_DEV_rproxy=$(DOCKER_COMPOSE_SERVICES_FOR_DEV_uisvr) uisvr
DOCKER_COMPOSE_SERVICES_FOR_DEV_all=$(DOCKER_COMPOSE_SERVICES_FOR_DEV_rproxy) rproxy

SERVICES?=$(DOCKER_COMPOSE_SERVICES_FOR_DEV_$(DEV_TARGET))

.PHONY: validate_services
validate_services:
	@if [ "$(DOCKER_COMPOSE_SERVICES_FOR_DEV_all)" = "$(DOCKER_COMPOSE_SERVICES)" ] ;\
	then \
		echo "services are valid"; \
	else \
		echo "services are invalid"; \
		exit 1; \
	fi

.PHONY: setup
setup: validate_services $(APP_RPROXY_ENVOY_YAML)
	$(MAKE) -C $(PATH_TO_UISVR) setup

.PHONY: run
run: setup
	$(ENVS) docker compose up ${SERVICES}

.PHONY: up
up: setup
	$(ENVS) docker compose up -d ${SERVICES}

.PHONY: down
down:
	$(ENVS) docker compose down ${SERVICES}


.PHONY: rmi
rmi: down
	for service in $(DOCKER_COMPOSE_SERVICES); do \
		docker rmi $(DOCKER_COMPOSE_NAME)-$$service:latest; \
	done

.PHONY: rebuild
rebuild:
	$(ENVS) $(MAKE) rmi ; \
	$(ENVS) docker compose build
