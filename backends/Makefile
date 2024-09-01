.PHONY: default
default: build lint test

PATH_TO_ROOT=..
include $(PATH_TO_ROOT)/.shapeapp.mk
include $(PATH_TO_SHAPEAPPMK)/make/eachdir.mk

DIRS=applib dbmigrations biz apisvr apisvr/proto

.PHONY: build
build:
	$(call eachdir,build,$(DIRS))

.PHONY: lint
lint:
	$(call eachdir,lint,$(DIRS))

.PHONY: test
test:
	$(call eachdir,test,$(DIRS))
