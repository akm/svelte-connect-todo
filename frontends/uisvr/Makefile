.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

.PHONY: install
install: npm-install connect-web-install

.PHONY: setup
setup: npm-ci

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk

.env:
	$(MAKE) dot_env_gen

.PHONY: dot_env_gen
dot_env_gen: $(TEXT_TEMPLATE_CLI)
	$(CONFIG_ENVS) \
	$(DEFAULT_PORT_ENVS) \
	$(TEXT_TEMPLATE_CLI) .env-template > .env

PATH_TO_DEV_CONTAINERS=$(PATH_TO_ROOT)/stages/local
.PHONY: dev-containers-up
dev-containers-up:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_DEV_CONTAINERS) up

.PHONY: dev-containers-down
dev-containers-down:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_DEV_CONTAINERS) down

.PHONY: dev_run
dev_run: setup .env
	npm run dev

.PHONY: dev
dev: dev-containers-up dev_run

include $(PATH_TO_SHAPEAPPMK)/sveltekit/npm.mk
include $(PATH_TO_SHAPEAPPMK)/sveltekit/app.mk

include $(PATH_TO_SHAPEAPPMK)/connect-web/install.mk
PATH_TO_PROTO=$(PATH_TO_APISVR)/proto
include $(PATH_TO_SHAPEAPPMK)/connect-web/generate.mk

.PHONY: generate
generate: connect-web-generate

DOCKER_IMAGE_REPOSITORY=$(APP_BASE_NAME)-uisvr
DOCKER_IMAGE_NAME?=$(DOCKER_IMAGE_REPOSITORY):$(APP_STAGE_TYPE)
include $(PATH_TO_SHAPEAPPMK)/docker/image.mk
