.PHONY: default
default: build lint test

PATH_TO_ROOT=.
include $(PATH_TO_ROOT)/.shapeapp.mk
include $(PATH_TO_SHAPEAPPMK)/default/setup.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/make/eachdir.mk

DIRS=backends frontends

.PHONY: build
build:
	$(call eachdir,build,$(DIRS))

.PHONY: lint
lint:
	$(call eachdir,lint,$(DIRS))

.PHONY: test
test:
	$(call eachdir,test,$(DIRS))
