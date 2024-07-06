.PHONY: default
default: build lint test

PATH_TO_ROOT=..
include $(PATH_TO_ROOT)/.shapeapp.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/frontends/children.mk

.PHONY: build
build:
	$(MAKE) -C uisvr build

.PHONY: lint
lint:
	$(MAKE) -C uisvr lint

.PHONY: test
test:
	$(MAKE) -C uisvr test_integration_pw_setup && \
	$(MAKE) -C uisvr test

.PHONY: dev
dev:
	$(MAKE) -C uisvr dev
