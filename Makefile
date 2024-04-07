.PHONY: default
default: build lint test

PATH_TO_PROJECT=.
include $(PATH_TO_PROJECT)/.govelte.config.mk
include $(PATH_TO_GOVELTEMK)/makefiles/root/asdf.mk
include $(PATH_TO_GOVELTEMK)/makefiles/root/children.mk

.PHONY: build
build:
	$(MAKE) -C backends build && \
	$(MAKE) -C frontends build

.PHONY: lint
lint:
	$(MAKE) -C backends lint && \
	$(MAKE) -C frontends lint

.PHONY: test
test:
	$(MAKE) -C backends test && \
	$(MAKE) -C frontends test

.PHONY: dev
dev:
	$(MAKE) -C frontends dev
