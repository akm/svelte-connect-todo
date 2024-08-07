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

PATH_TO_LOCAL=$(PATH_TO_ROOT)/stages/local
.PHONY: dev-containers-up
dev-containers-up:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_LOCAL) up

.PHONY: dev-containers-down
dev-containers-down:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_LOCAL) down

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

APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-uisvr
APP_CONTAINER_IMAGE_LOCAL=$(APP_CONTAINER_IMAGE_NAME):local

.PHONY: container_image_name
container_image_name:
	@echo $(APP_CONTAINER_IMAGE_NAME)

.PHONY: container_build
container_build:
	docker build \
		-t $(APP_CONTAINER_IMAGE_LOCAL) \
		-f Dockerfile \
		.
