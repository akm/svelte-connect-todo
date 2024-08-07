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
	DB_DSN='$(shell DEV_TARGET=apisvr $(MAKE) -C $(PATH_TO_LOCAL) mysql-dsn)' \
	$(foreach var,$(DEV_VARS),$(var)=$($(var)))

PATH_TO_CMD_APISVR=./cmd/server
.PHONY: run
run:
	$(DEV_ENVS) go run $(PATH_TO_CMD_APISVR)

.PHONY: dev
dev: dev_container_up run

GOLANG_BUILD_TARGET=./cmd/server
GOLANG_BUILD_OUTPUT_BASE=./bin/server

GOLANG_BUILD_OUTPUT_linux_amd64=$(GOLANG_BUILD_OUTPUT_BASE)-linux-amd64
.PHONY: golang-build-linux-amd64
golang-build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o $(GOLANG_BUILD_OUTPUT_linux_amd64) $(GOLANG_BUILD_TARGET)

GOLANG_BUILD_OUTPUT_linux_arm64=$(GOLANG_BUILD_OUTPUT_BASE)-linux-arm64
.PHONY: golang-build-linux-arm64
golang-build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
		go build -o $(GOLANG_BUILD_OUTPUT_linux_arm64) $(GOLANG_BUILD_TARGET)

GOLANG_BUILD_OUTPUT_darwin_arm64=$(GOLANG_BUILD_OUTPUT_BASE)-darwin-arm64
.PHONY: golang-build-darwin-arm64
golang-build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
		go build -o $(GOLANG_BUILD_OUTPUT_darwin_arm64) $(GOLANG_BUILD_TARGET)

GOLANG_BUILD_OUTPUT_dev=$(GOLANG_BUILD_OUTPUT_BASE)-dev
.PHONY: golang-build-dev
golang-build-dev:
	go build -o $(GOLANG_BUILD_OUTPUT_dev) $(GOLANG_BUILD_TARGET)

GOLANG_BUILD_OUTPUT_FOR_local=$(GOLANG_BUILD_OUTPUT_linux_$(LOCAL_GOARCH))
$(GOLANG_BUILD_OUTPUT_FOR_local):
	$(MAKE) build-local
.PHONY: build-local
build-local: golang-build-linux-$(LOCAL_GOARCH)
.PHONY: binary-file-for-local
binary-file-for-local:
	@echo "$(GOLANG_BUILD_OUTPUT_FOR_local)"

GOLANG_BUILD_OUTPUT_FOR_production=$(GOLANG_BUILD_OUTPUT_linux_amd64)
$(GOLANG_BUILD_OUTPUT_FOR_production):
	$(MAKE) build-production
.PHONY: build-production
build-production: golang-build-linux-amd64
.PHONY: binary-file-for-production
binary-file-for-production:
	@echo "$(GOLANG_BUILD_OUTPUT_FOR_production)"

BINARY_FILE=$(GOLANG_BUILD_OUTPUT_FOR_$(APP_STAGE_TYPE))

include $(PATH_TO_SHAPEAPPMK)/golang/module.mk
APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-$(GOLANG_MODULE_NAME)
APP_CONTAINER_IMAGE_dev=$(APP_CONTAINER_IMAGE_NAME):dev

.PHONY: container_image_name
container_image_name:
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

PATH_TO_BACKENDS_TEST_CONTAINERS=$(PATH_TO_BACKENDS)/test/containers
.PHONY: test_container_up
test_container_up:
	$(MAKE) -C $(PATH_TO_BACKENDS_TEST_CONTAINERS) up

.PHONY: test_container_down
test_container_down:
	$(MAKE) -C $(PATH_TO_BACKENDS_TEST_CONTAINERS) down

# テストに渡される TEST_FIXTURE_DIR は絶対パス。テストはそれぞれの_test.goファイルのディレクトリで実行されるため。
TEST_FIXTURE_DIR=$(abspath $(PATH_TO_BIZ)/test/fixtures)

TEST_VARS=\
	APP_STAGE \
	APP_CORS_ALLOW_ORIGINS \
	APP_FIREBASE_API_KEY \
	FIREBASE_AUTH_EMULATOR_HOST
TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT_LOCAL) \
	TEST_DB_DSN='$(shell $(MAKE) -C $(PATH_TO_BACKENDS_TEST_CONTAINERS) mysql-dsn)' \
	TEST_FIXTURE_DIR=$(TEST_FIXTURE_DIR) \
	$(foreach var,$(TEST_VARS),$(var)=$($(var)))

.PHONY: test
test: test_container_up
	$(TEST_ENVS) go test ./...
