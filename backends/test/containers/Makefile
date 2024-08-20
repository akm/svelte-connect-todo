PATH_TO_ROOT=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/docker-compose/base.mk

TEST_MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME)-test

DOCKER_COMPOSE_ENVS=\
	$(DEFAULT_PORT_ENVS) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	APP_MYSQL_DB_NAME=$(TEST_MYSQL_DB_NAME)

DOCKER_COMPOSE_SERVICE_REQUIRED?=mysql

TARGET_SERVICES_WITH_mysql=mysql
TARGET_SERVICES_WITH_firebase_emulators=firebase-emulators $(TARGET_SERVICES_WITH_mysql)
TARGET_SERVICES_WITH_ALL=$(TARGET_SERVICES_WITH_firebase_emulators)
DOCKER_COMPOSE_TARGET_SERVICES?=$(TARGET_SERVICES_WITH_$(DOCKER_COMPOSE_SERVICE_REQUIRED))

ifneq "$(sort $(TARGET_SERVICES_WITH_ALL))" "$(sort $(DOCKER_COMPOSE_SERVICES_ALL))"
$(error "services are invalid")
endif

.PHONY: run
run: docker-compose-run
.PHONY: up
up: docker-compose-up mysql-wait dbmigration-up
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
	MYSQL_DB_NAME=$(TEST_MYSQL_DB_NAME)

include $(PATH_TO_SHAPEAPPMK)/mysql/base.mk

.PHONY: dbmigration-up
dbmigration-up:
	APP_ENV=unit_test GOOSE_DSN='$(MYSQL_DSN)' \
	$(MAKE) -C $(PATH_TO_DBMIGRATIONS) up
