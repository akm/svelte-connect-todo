.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/components/organelles/golang/base.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/git/check.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/sqlc/base.mk

include $(PATH_TO_SHAPEAPPMK)/components/organs/biz/schema-sql.mk
include $(PATH_TO_SHAPEAPPMK)/components/organs/biz/setup.mk

.PHONY: test
test:
	go test ./...

SQLC_CLI_VERSION=v1.26.0
