.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.config.mk

GO_ROOT_PACKAGE=apisvr

include $(PATH_TO_SHAPEAPPMK)/makefiles/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/git/check.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/golang/build.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/golang/lint.mk

# connect-go
include $(PATH_TO_SHAPEAPPMK)/makefiles/apisvr/connect-go.mk

# run, run_with_devbug

# DEV_MYSQL_DSN?='$(shell $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)'

# DEV_SERVER_ENVS_BASE=\
# 	APP_STAGE=$(APP_STAGE) \
# 	APP_MYSQL_DSN=$(DEV_MYSQL_DSN) \
# 	APP_LOG_CONSOLE_WRITER=true \
# 	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

DEV_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

# dev, dev_container_up, dev_container_down
# PATH_TO_CMD_APISVR needs dot as an argument fo go run
PATH_TO_CMD_APISVR=./cmd/server
include $(PATH_TO_SHAPEAPPMK)/makefiles/apisvr/dev.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_SHAPEAPPMK)/makefiles/apisvr/test.mk

APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-$(GO_ROOT_PACKAGE)
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

.PHONY: test-connections
test-connections:
ifeq ($(SERVER_DOMAIN),)
	@echo "SERVER_DOMAIN is not set. Please set SERVER_DOMAIN to test connections."
else
	$(MAKE) install_tools && \
	$(MAKE) -C proto build && \
	./test-connections.sh
endif
