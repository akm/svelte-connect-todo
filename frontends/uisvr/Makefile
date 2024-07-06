.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.config.mk

.PHONY: setup
setup: install

DEV_ENVS=\
	VITE_GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	VITE_FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	VITE_UISVR_ORIGIN=http://localhost:$(APP_PORT_UISVR_dev)

BUILD_DEPS=setup

include $(PATH_TO_SHAPEAPPMK)/makefiles/uisvr/base.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/uisvr/dev.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/uisvr/connect-web.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/uisvr/test_integration.mk

node_run: build
	PORT=$(APP_PORT_UISVR_e2e_test) node -r dotenv/config build

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
