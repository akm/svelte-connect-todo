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

# DEV_CONTAINERS-mysql-dsn
$(call shell-dir-target-vars,DEV_CONTAINERS-,$(PATH_TO_ROOT)/stages/local,mysql-dsn)

APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test)
FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/env.mk

APISVR_DEV_VARS=$(APISVR_ENV_VARS)
APISVR_DEV_ENVS=$(APISVR_ENVS_BASE) \
	DB_DSN='$(DEV_CONTAINERS-mysql-dsn)' \
	$(foreach var,$(APISVR_DEV_VARS),$(var)=$($(var)))

APISVR_DEV_TARGET_PACKAGE=./cmd/server
APISVR_DEV_PATH_TO_CONTAINERS=$(PATH_TO_ROOT)/stages/local
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk

.PHONY: run
run: dev-run

.PHONY: dev
dev: dev-containers-up run

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

# TEST_CONTAINERS-mysql-dsn
$(call shell-dir-target-vars,TEST_CONTAINERS-,$(PATH_TO_BACKENDS)/test/containers,mysql-dsn)


# テストに渡される TEST_FIXTURE_DIR は絶対パス。テストはそれぞれの_test.goファイルのディレクトリで実行されるため。
APISVR_TEST_FIXTURE_DIR=$(abspath $(PATH_TO_BIZ)/test/fixtures)

APISVR_TEST_VARS=$(APISVR_ENV_VARS)
APISVR_TEST_ENVS=$(APISVR_ENVS_BASE) \
	DB_DSN='$(TEST_CONTAINERS-mysql-dsn)' \
	TEST_FIXTURE_DIR=$(APISVR_TEST_FIXTURE_DIR) \
	$(foreach var,$(APISVR_TEST_VARS),$(var)=$($(var)))

APISVR_TEST_PATH_TO_CONTAINERS=$(PATH_TO_BACKENDS)/test/containers
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk

.PHONY: test
test: test-containers-up test-run
