.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

GO_ROOT_PACKAGE=apisvr

include $(PATH_TO_GOVELTEMK)/makefiles/asdf/reshim.mk
include $(PATH_TO_GOVELTEMK)/makefiles/git/check.mk
include $(PATH_TO_GOVELTEMK)/makefiles/golang/build.mk
include $(PATH_TO_GOVELTEMK)/makefiles/golang/lint.mk

# connect-go
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/connect-go.mk

# run, run_with_devbug
# PATH_TO_CMD_APISVR needs dot as an argument fo go run
PATH_TO_CMD_APISVR=./cmd/server
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/run.mk

# dev, dev_container_up, dev_container_down
DEV_SERVER_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/dev.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/test.mk
