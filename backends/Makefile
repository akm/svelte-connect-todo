.PHONY: default
default: build lint test

PATH_TO_PROJECT=..
include $(PATH_TO_PROJECT)/.govelte.config.mk
include $(PATH_TO_GOVELTEMK)/makefiles/backends/children.mk

.PHONY: build
build:
	$(MAKE) -C applib build && \
	$(MAKE) -C biz build && \
	$(MAKE) -C apisvr build

.PHONY: lint
lint:
	$(MAKE) -C applib lint && \
	$(MAKE) -C biz lint && \
	$(MAKE) -C apisvr lint

.PHONY: test
test:
	$(MAKE) -C applib test && \
	$(MAKE) -C biz test && \
	$(MAKE) -C apisvr test

.PHONY: dev
dev:
	$(MAKE) -C containers dev && \
	$(MAKE) -C dbmigrations dev && \
	$(MAKE) -C apisvr dev
