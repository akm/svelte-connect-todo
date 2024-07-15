.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

GOOSE_TARGET_STAGE?=local/dev
GOOSE_TARGET_STAGE_PATH?=$(PATH_TO_ROOT)/stages/$(GOOSE_TARGET_STAGE)
GOOSE_DSN=$(shell $(MAKE) -C $(GOOSE_TARGET_STAGE_PATH) mysql-dsn)
GOOSE_MAIN_PACKAGE=./cmd/runner
GOOSE_ENVS=\
	APP_ENV=$(APP_ENV) \
	APP_STAGE=$(APP_STAGE) \
	APP_STAGE_TYPE=$(APP_STAGE_TYPE)
include $(PATH_TO_SHAPEAPPMK)/goose/commands.mk
include $(PATH_TO_SHAPEAPPMK)/goose/app.mk
