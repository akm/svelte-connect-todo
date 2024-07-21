.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk

include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
build: golang-build

include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
lint: golangci-lint-lint

include $(PATH_TO_SHAPEAPPMK)/sqlc/base.mk

SCHEMA_SQL=schema.sql
$(SCHEMA_SQL): dump-schema-sql

MYSQL_ENVS=$(shell $(MAKE) -C ../test/containers mysql-envs --no-print-directory)

.PHONY: test-containers-up
test-containers-up:
	$(MAKE) -C ../test/containers up

.PHONY: test-containers-down
test-containers-down:
	$(MAKE) -C ../test/containers down

.PHONY: dump-schema-sql
dump-schema-sql: test-containers-down test-containers-up dump-schema-sql-impl
	$(MAKE) test-containers-down

.PHONY: dump-schema-sql-impl
dump-schema-sql-impl:
	$(MYSQL_ENVS) \
	DUMP_OPTS=--no-data \
	DUMP_POST_PROCESS='| grep -v "Dump completed on" > $(abspath $(SCHEMA_SQL))' \
	$(MAKE) -C $(PATH_TO_SHAPEAPPMK)/mysql dump
