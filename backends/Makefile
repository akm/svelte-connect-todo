.PHONY: default
default: build lint test

PATH_TO_ROOT=..
include $(PATH_TO_ROOT)/.shapeapp.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/backends/children.mk

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
	$(MAKE) -C apisvr dev
