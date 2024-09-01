.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/golang/base.mk
include $(PATH_TO_SHAPEAPPMK)/git/check.mk
include $(PATH_TO_SHAPEAPPMK)/sqlc/base.mk

include $(PATH_TO_SHAPEAPPMK)/default/biz/schema-sql.mk

.PHONY: test
test:
	go test ./...

SQLC_CLI_VERSION=v1.26.0
