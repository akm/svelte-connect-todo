.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

.PHONY: setup
setup: install
	$(MAKE) $(FIREBASE_CONFIG_FILE)

FIREBASE_CONFIG_FILE=src/lib/firebase/firebaseconfig.ts
$(FIREBASE_CONFIG_FILE):
	cp $(FIREBASE_CONFIG_FILE).dummy $(FIREBASE_CONFIG_FILE)

DEV_ENVS=\
	VITE_GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	VITE_FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	VITE_UISVR_ORIGIN=http://localhost:$(APP_PORT_UISVR_dev)

BUILD_DEPS=setup

include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/base.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/dev.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/connect-web.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/test_integration.mk

node_run: build
	PORT=$(APP_PORT_UISVR_e2e_test) node -r dotenv/config build

APP_CONTAINER_REPOSITORY?=akima
APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-uisvr
APP_CONTAINER_IMAGE_TAG=$(shell git show --format='%h' --no-patch)
APP_CONTAINER_IMAGE_LOCAL=$(APP_CONTAINER_IMAGE_NAME):local
APP_CONTAINER_IMAGE_PRODUCTION=$(APP_CONTAINER_IMAGE_NAME):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_image_name
container_image_name:
	@echo $(APP_CONTAINER_IMAGE)

.PHONY: container_build
container_build: container_build_local

.PHONY: container_build_local
container_build_local:
	docker build \
		-t $(APP_CONTAINER_IMAGE_LOCAL) \
		-f Dockerfile \
		.

.PHONY: container_build_production
container_build_production:
	docker build \
		--platform linux/amd64 \
		-t $(APP_CONTAINER_IMAGE_PRODUCTION) \
		-f Dockerfile \
		.

.PHONY: container_push
container_push: container_build_production
	docker tag $(APP_CONTAINER_IMAGE_PRODUCTION) $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) && \
	docker push $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION)
