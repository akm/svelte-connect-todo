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
	$(shell $(MAKE) default-port-envs --no-print-directory) \
	$(TEXT_TEMPLATE_CLI) .env-template > .env

PATH_TO_LOCAL_DEV=$(PATH_TO_ROOT)/stages/local/dev
.PHONY: dev_container_up
dev_container_up:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_LOCAL_DEV) up

.PHONY: dev_container_down
dev_container_down:
	DEV_TARGET=uisvr $(MAKE) -C $(PATH_TO_LOCAL_DEV) down

.PHONY: dev_run
dev_run: setup .env
	npm run dev

.PHONY: dev
dev: dev_container_up dev_run

BUILD_DEPS=setup
include $(PATH_TO_SHAPEAPPMK)/sveltekit/npm.mk

include $(PATH_TO_SHAPEAPPMK)/connect-web/install.mk
include $(PATH_TO_SHAPEAPPMK)/connect-web/generate.mk

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
