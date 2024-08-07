.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build-all

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

PATH_TO_PROTO=./proto
include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/app/gen.mk

.PHONY: generate
generate: buf-generate

PATH_TO_DEV_CONTAINERS=$(PATH_TO_ROOT)/stages/local
APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test)
FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

ENVS_BASE=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL)
ENV_VARS=\
	APP_STAGE \
	APP_CORS_ALLOW_ORIGINS \
	APP_FIREBASE_API_KEY \
	FIREBASE_AUTH_EMULATOR_HOST

DEV_VARS=$(ENV_VARS)
DEV_ENVS=$(ENVS_BASE) \
	DB_DSN='$(shell DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_DEV_CONTAINERS) mysql-dsn)' \
	$(foreach var,$(DEV_VARS),$(var)=$($(var)))

PATH_TO_DEV_RUN_PACKAGE=./cmd/server
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk

.PHONY: run
run: dev-run

.PHONY: dev
dev: dev-containers-up run

GOLANG_BINARY_TARGET=./cmd/server
GOLANG_BINARY_OUTPUT_BASE=./bin/server
include $(PATH_TO_SHAPEAPPMK)/golang/binary.mk

BINARY_FILE=$(GOLANG_BINARY_PATH_FOR_STAGE_$(APP_STAGE_TYPE))

include $(PATH_TO_SHAPEAPPMK)/golang/module.mk
APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-$(GOLANG_MODULE_NAME)
APP_CONTAINER_IMAGE_dev=$(APP_CONTAINER_IMAGE_NAME):dev

.PHONY: docker-container-image-name
docker-container-image-name:
	@echo $(APP_CONTAINER_IMAGE_NAME)

.PHONY: container_build
container_build: $(BINARY_FILE)
	docker build \
		-t $(APP_CONTAINER_IMAGE_dev) \
		--build-arg APP_BIN_PATH=$(BINARY_FILE) \
		-f Dockerfile \
		.

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk

.PHONY: test-connections
test-connections: 
	$(MAKE) -C proto test-connections

PATH_TO_TEST_CONTAINERS=$(PATH_TO_BACKENDS)/test/containers

# テストに渡される TEST_FIXTURE_DIR は絶対パス。テストはそれぞれの_test.goファイルのディレクトリで実行されるため。
TEST_FIXTURE_DIR=$(abspath $(PATH_TO_BIZ)/test/fixtures)

TEST_VARS=$(ENV_VARS) \
	TEST_FIXTURE_DIR
TEST_ENVS=$(ENVS_BASE) \
	TEST_DB_DSN='$(shell $(MAKE) -C $(PATH_TO_TEST_CONTAINERS) mysql-dsn)' \
	$(foreach var,$(TEST_VARS),$(var)=$($(var)))

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk

.PHONY: test
test: test-containers-up test-run
