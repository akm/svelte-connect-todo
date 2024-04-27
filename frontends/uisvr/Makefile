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
	VITE_FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	VITE_UISVR_ORIGIN=http://localhost:$(APP_PORT_UISVR_HTTP_dev)

include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/base.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/dev.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/connect-web.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/test_integration.mk

build:
	npm run build

node_run: build
	PORT=4173 node -r dotenv/config build
