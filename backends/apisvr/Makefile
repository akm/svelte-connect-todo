.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/golang/build.mk
include $(PATH_TO_SHAPEAPPMK)/golangci-lint/lint.mk
build: golang-build-all
lint: golangci-lint-lint

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/buf.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/env.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/docker-image.mk

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk
.PHONY: test-connections
test-connections:
	$(MAKE) -C proto test-connections
