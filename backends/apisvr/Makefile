.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk

PATH_TO_PROTO=./proto
include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/app/gen.mk

PATH_TO_LOCAL_DEV=$(PATH_TO_ROOT)/stages/local/dev
.PHONY: dev_container_up
dev_container_up:
	DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_LOCAL_DEV) up

.PHONY: dev_container_down
dev_container_down:
	DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_LOCAL_DEV) down

DEV_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

PATH_TO_CMD_APISVR=./cmd/server
.PHONY: run
run:
	$(DEV_ENVS) go run $(PATH_TO_CMD_APISVR)

.PHONY: dev
dev: dev_container_up run


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

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk

.PHONY: test-connections
test-connections: $(GRPCURL_CLI) $(BUF_CLI)
	$(MAKE) -C proto build && \
	./test-connections.sh
