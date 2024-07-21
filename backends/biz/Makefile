.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

SCHEMA_SQL=schema.sql
$(SCHEMA_SQL): dump-schema-sql

MYSQL_ENVS=$(shell $(MAKE) -C ../test/containers mysql-envs --no-print-directory)

.PHONY: dump-schema-sql
dump-schema-sql:
	$(MYSQL_ENVS) \
	DUMP_OPTS=--no-data \
	DUMP_DEST=$(abspath $(SCHEMA_SQL)) \
	$(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql dump
