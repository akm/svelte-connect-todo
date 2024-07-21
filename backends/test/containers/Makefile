PATH_TO_ROOT=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk
include $(PATH_TO_SHAPEAPPMK)/docker-compose/base.mk

DOCKER_COMPOSE_ENVS=\
	$(DEFAULT_PORT_ENVS) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)

.PHONY: docker_compose_envs
docker_compose_envs:
	@echo $(DOCKER_COMPOSE_ENVS)

DOCKER_COMPOSE_SERVICE_REQUIRED?=mysql

TARGET_SERVICES_WITH_mysql=mysql
TARGET_SERVICES_WITH_firebase_emulators=firebase-emulators $(TARGET_SERVICES_WITH_mysql)
TARGET_SERVICES_WITH_ALL=$(TARGET_SERVICES_WITH_firebase_emulators)
TARGET_SERVICES?=$(TARGET_SERVICES_WITH_$(DOCKER_COMPOSE_SERVICE_REQUIRED))

.PHONY: validate_services
validate_services:
	@if [ "$(sort $(TARGET_SERVICES_WITH_ALL))" = "$(sort $(DOCKER_COMPOSE_SERVICES_ALL))" ] ;\
	then \
		echo "services are valid"; \
	else \
		echo "services are invalid"; \
		exit 1; \
	fi

.PHONY: setup
setup: validate_services

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
	MYSQL_PORT=$(APP_PORT_MYSQL_unit_test) \
	MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)

.PHONY: mysql-dsn
mysql-dsn:
	@$(MYSQL_ENVS) $(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql dsn --no-print-directory

.PHONY: mysql-wait
mysql-wait:
	$(MYSQL_ENVS) $(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql wait

.PHONY: mysql-console
mysql-console:
	$(MYSQL_ENVS) $(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql console
