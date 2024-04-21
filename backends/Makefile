.PHONY: default
default: build lint test

PATH_TO_PROJECT=..
include $(PATH_TO_PROJECT)/.govelte.config.mk
include $(PATH_TO_GOVELTEMK)/makefiles/backends/children.mk

# DIRECTORIES=applib biz apisvr
DIRECTORIES=apisvr

.PHONY: build
build:
	for dir in $(DIRECTORIES); do \
		$(MAKE) -C $$dir build; \
	done

.PHONY: lint
lint:
	for dir in $(DIRECTORIES); do \
		$(MAKE) -C $$dir lint; \
	done

.PHONY: test
test:
	for dir in $(DIRECTORIES); do \
		$(MAKE) -C $$dir test; \
	done

.PHONY: dev
dev:
	$(MAKE) -C apisvr dev; \
