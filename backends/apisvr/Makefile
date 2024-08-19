.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build-all

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

PATH_TO_PROTO=./proto
include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/app/gen.mk

.PHONY: generate
generate: buf-generate

APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test)
FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

APISVR_DEV_TARGET_PACKAGE?=./cmd/server
APISVR_DEV_PATH_TO_CONTAINERS?=$(PATH_TO_ROOT)/stages/local

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/env.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk

GOLANG_BINARY_TARGET=./cmd/server
GOLANG_BINARY_OUTPUT_BASE=./bin/server

DOCKER_IMAGE_BUILD_DEPS=$(BINARY_FILE)
DOCKER_IMAGE_BUILD_OPTS=\
	--build-arg APP_BIN_PATH=$(BINARY_FILE) \
	-f Dockerfile
include $(PATH_TO_SHAPEAPPMK)/golang/docker-image.mk

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk

.PHONY: test-connections
test-connections:
	$(MAKE) -C proto test-connections

# テストに渡される TEST_PATH_TO_FIXTURES は絶対パス。テストはそれぞれの_test.goファイルのディレクトリで実行されるため。
APISVR_TEST_PATH_TO_FIXTURES=$(abspath $(PATH_TO_BIZ)/test/fixtures)
APISVR_TEST_PATH_TO_CONTAINERS=$(PATH_TO_BACKENDS)/test/containers

# TEST_CONTAINERS-mysql-dsn
$(call shell-dir-target-vars,TEST_CONTAINERS-,$(APISVR_TEST_PATH_TO_CONTAINERS),mysql-dsn)

APISVR_TEST_VARS=$(APISVR_ENV_VARS)
APISVR_TEST_ENVS=$(APISVR_ENVS_BASE) \
	DB_DSN='$(TEST_CONTAINERS-mysql-dsn)' \
	TEST_PATH_TO_FIXTURES=$(APISVR_TEST_PATH_TO_FIXTURES) \
	$(foreach var,$(APISVR_TEST_VARS),$(var)=$($(var)))

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk

.PHONY: test
test: test-containers-up test-run
