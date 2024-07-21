PATH_TO_ROOT=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk

ENVS=\
	$(DEFAULT_PORT_ENVS) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)

.PHONY: envs
envs:
	@echo $(ENVS)

DOCKER_COMPOSE_NAME=$(shell cat ./docker-compose.yml | yq .name)
DOCKER_COMPOSE_SERVICES_FULL=$(shell cat docker-compose.yml | yq '.services | select(.build.context != "") | keys | join(" ")')

DOCKER_COMPOSE_SERVICE_REQUIRED?=mysql

DOCKER_COMPOSE_SERVICES_WITH_mysql=mysql
DOCKER_COMPOSE_SERVICES_WITH_firebase_emulators=firebase-emulators $(DOCKER_COMPOSE_SERVICES_WITH_mysql)
DOCKER_COMPOSE_SERVICES_all=$(DOCKER_COMPOSE_SERVICES_WITH_firebase_emulators)

SERVICES?=$(DOCKER_COMPOSE_SERVICES_WITH_$(DOCKER_COMPOSE_SERVICE_REQUIRED))

.PHONY: validate_services
validate_services:
	@if [ "$(DOCKER_COMPOSE_SERVICES_all)" = "$(DOCKER_COMPOSE_SERVICES_FULL)" ] ;\
	then \
		echo "services are valid"; \
	else \
		echo "services are invalid"; \
		exit 1; \
	fi

.PHONY: setup
setup: validate_services

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
	for service in $(DOCKER_COMPOSE_SERVICES_FULL); do \
		docker rmi $(DOCKER_COMPOSE_NAME)-$$service:latest; \
	done

.PHONY: rebuild
rebuild:
	$(ENVS) $(MAKE) rmi ; \
	$(ENVS) docker compose build

MYSQL_ENVS=\
	MYSQL_USER_NAME=root \
	MYSQL_USER_PASSWORD= \
	MYSQL_HOST=127.0.0.1 \
	MYSQL_PORT=$(APP_PORT_MYSQL_unit_test) \
	MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)

.PHONY: mysql-dsn
mysql-dsn:
	@$(MYSQL_ENVS) $(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql dsn --no-print-directory

.PHONY: mysql-console
mysql-console:
	$(MYSQL_ENVS) $(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql console
