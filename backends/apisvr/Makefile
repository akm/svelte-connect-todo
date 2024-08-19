.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/default/apisvr/golang.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/buf.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/env.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/dev.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/test.mk
include $(PATH_TO_SHAPEAPPMK)/default/apisvr/docker-image.mk

.PHONY: test-connections
test-connections:
	$(MAKE) -C proto test-connections
