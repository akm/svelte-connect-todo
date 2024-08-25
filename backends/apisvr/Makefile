.PHONY: default
default: build lint test

PATH_TO_ROOT:=../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/default/apisvr.mk

.PHONY: test-connections
test-connections:
	$(MAKE) -C proto test-connections
