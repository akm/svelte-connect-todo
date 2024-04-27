.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

.PHONY: setup
setup: $(NODE_MODULES) 
	$(MAKE) $(FIREBASE_CONFIG_FILE)

FIREBASE_CONFIG_FILE=src/lib/firebase/firebaseconfig.ts
$(FIREBASE_CONFIG_FILE):
	cp $(FIREBASE_CONFIG_FILE).dummy $(FIREBASE_CONFIG_FILE)

DEV_ENVS=\
	VITE_GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	VITE_FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/base.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/dev.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/connect-web.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/test_integration.mk

APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-uisvr
APP_CONTAINER_IMAGE_TAG=$(shell git show --format='%h' --no-patch)
APP_CONTAINER_IMAGE=$(APP_CONTAINER_IMAGE_NAME):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_build
container_build:
	$(DEV_ENVS) docker build . -t $(APP_CONTAINER_IMAGE)
