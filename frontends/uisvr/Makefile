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

include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/base.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/dev.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/connect-web.mk
include $(PATH_TO_GOVELTEMK)/makefiles/uisvr/test_integration.mk
