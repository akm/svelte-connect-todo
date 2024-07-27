.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

PATH_TO_PROTO=./proto
include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/app/gen.mk

.PHONY: generate
generate: buf-generate

PATH_TO_LOCAL=$(PATH_TO_ROOT)/stages/local
.PHONY: dev_container_up
dev_container_up:
	DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_LOCAL) up

.PHONY: dev_container_down
dev_container_down:
	DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_LOCAL) down

APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test)
FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

DEV_VARS=\
	APP_STAGE \
	APP_CORS_ALLOW_ORIGINS \
	APP_FIREBASE_API_KEY \
	FIREBASE_AUTH_EMULATOR_HOST
DEV_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	$(foreach var,$(DEV_VARS),$(var)=$($(var)))

PATH_TO_CMD_APISVR=./cmd/server
.PHONY: run
run:
	$(DEV_ENVS) go run $(PATH_TO_CMD_APISVR)

.PHONY: dev
dev: dev_container_up run

PATRH_TO_BIN_LOCAL=bin/server-local
.PHONY: path_to_bin_local
path_to_bin_local:
	@echo $(PATRH_TO_BIN_LOCAL)
$(PATRH_TO_BIN_LOCAL): build_bin_local
.PHONY: build_bin_local
build_bin_local:
	go build -o $(PATRH_TO_BIN_LOCAL) $(PATH_TO_CMD_APISVR)

PATRH_TO_BIN_PRODUCTION=bin/server-production
.PHONY: path_to_bin_production
path_to_bin_production:
	@echo $(PATRH_TO_BIN_PRODUCTION)
$(PATRH_TO_BIN_PRODUCTION): build_bin_production
.PHONY: build_bin_production
build_bin_production:
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go mod download && \
	go build -mod=readonly -v -o $(PATRH_TO_BIN_PRODUCTION) $(PATH_TO_CMD_APISVR)

include $(PATH_TO_SHAPEAPPMK)/golang/module.mk
APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-$(GOLANG_MODULE_NAME)
APP_CONTAINER_IMAGE_LOCAL=$(APP_CONTAINER_IMAGE_NAME):local

.PHONY: container_image_name
container_image_name:
	@echo $(APP_CONTAINER_IMAGE_NAME)

.PHONY: container_build
container_build: build_bin_local
	docker build \
		-t $(APP_CONTAINER_IMAGE_LOCAL) \
		-f Dockerfile \
		.

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk

.PHONY: test-connections
test-connections: $(GRPCURL_CLI) $(BUF_CLI)
	$(MAKE) -C proto build && \
	./test-connections.sh
