.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/golang/base.mk

GOOSE_TARGET_ENV_PATH?=$(PATH_TO_ROOT)/stages/local
GOOSE_DSN?=$(shell $(MAKE) -C $(GOOSE_TARGET_ENV_PATH) mysql-dsn)
GOOSE_MAIN_PACKAGE=./cmd/runner
GOOSE_ENVS=\
	APP_ENV=$(APP_ENV) \
	APP_STAGE=$(APP_STAGE) \
	APP_STAGE_TYPE=$(APP_STAGE_TYPE)
include $(PATH_TO_SHAPEAPPMK)/goose/commands.mk
include $(PATH_TO_SHAPEAPPMK)/default/dbmigrations.mk
