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

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/env.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk

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

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk
