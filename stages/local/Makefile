PATH_TO_ROOT=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk
include $(PATH_TO_SHAPEAPPMK)/docker-compose/base.mk

DEV_TARGET?=all

DOCKER_COMPOSE_ENVS_BASE=\
	$(DEFAULT_PORT_ENVS) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME) \
	DEV_TARGET=$(DEV_TARGET)

APP_UISVR_DOT_ENV_PATH=./uisvr.env
$(APP_UISVR_DOT_ENV_PATH): uisvr_env_gen
.PHONY: uisvr_env_gen
uisvr_env_gen: $(TEXT_TEMPLATE_CLI)
	$(DOCKER_COMPOSE_ENVS_BASE) $(TEXT_TEMPLATE_CLI) ./uisvr.env.template > $(APP_UISVR_DOT_ENV_PATH)

APP_RPROXY_ENVOY_YAML_TEMPLATE=$(PATH_TO_MIDDLEWARES)/rproxy/local/envoy.yaml.template
APP_RPROXY_ENVOY_YAML=./envoy.yaml
$(APP_RPROXY_ENVOY_YAML): envoy_yaml_gen
.PHONY: envoy_yaml_gen
envoy_yaml_gen: $(TEXT_TEMPLATE_CLI)
	$(DOCKER_COMPOSE_ENVS_BASE) $(TEXT_TEMPLATE_CLI) $(APP_RPROXY_ENVOY_YAML_TEMPLATE) > $(APP_RPROXY_ENVOY_YAML)

DOCKER_COMPOSE_ENVS=\
	$(DOCKER_COMPOSE_ENVS_BASE) \
	APP_UISVR_DOT_ENV=$(APP_UISVR_DOT_ENV_PATH) \
	APP_RPROXY_ENVOY_YAML=$(APP_RPROXY_ENVOY_YAML)

.PHONY: docker_compose_envs
docker_compose_envs:
	@echo $(DOCKER_COMPOSE_ENVS)

.PHONY: setup_envs
setup_envs: $(APP_UISVR_DOT_ENV_PATH) $(APP_RPROXY_ENVOY_YAML)

TARGET_SERVICES_FOR_DEV_apisvr=firebase-emulators mysql
TARGET_SERVICES_FOR_DEV_uisvr=$(TARGET_SERVICES_FOR_DEV_apisvr) apisvr
TARGET_SERVICES_FOR_DEV_rproxy=$(TARGET_SERVICES_FOR_DEV_uisvr) uisvr
TARGET_SERVICES_FOR_DEV_all=$(TARGET_SERVICES_FOR_DEV_rproxy) rproxy
DOCKER_COMPOSE_TARGET_SERVICES?=$(TARGET_SERVICES_FOR_DEV_$(DEV_TARGET))

.PHONY: validate_services
validate_services:
	if [ "$(sort $(TARGET_SERVICES_FOR_DEV_all))" = "$(sort $(DOCKER_COMPOSE_SERVICES_ALL))" ] ;\
	then \
		echo "services are valid"; \
	else \
		echo "services are invalid"; \
		exit 1; \
	fi

.PHONY: setup
setup: validate_services setup_envs
	$(MAKE) -C $(PATH_TO_UISVR) setup

.PHONY: run
run: setup docker-compose-run
.PHONY: up
up: setup docker-compose-up
.PHONY: down
down: docker-compose-down
.PHONY: rmi
rmi: docker-compose-rmi
.PHONY: rebuild
rebuild: docker-compose-rebuild

MYSQL_ENVS=\
	MYSQL_USER_NAME=root \
	MYSQL_USER_PASSWORD= \
	MYSQL_HOST=127.0.0.1 \
	MYSQL_PORT=$(APP_PORT_MYSQL_dev) \
	MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)

include $(PATH_TO_SHAPEAPPMK)/mysql/base.mk

.PHONY: dbmigration-up
dbmigration-up:
	GOOSE_DSN='$(MYSQL_DSN)' \
	$(MAKE) -C $(PATH_TO_DBMIGRATIONS) up
