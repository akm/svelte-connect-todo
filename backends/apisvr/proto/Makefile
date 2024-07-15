.PHONY: default
default: build lint test

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/proto/base.mk
include $(PATH_TO_SHAPEAPPMK)/protobuf/protoc-gen-go.mk
include $(PATH_TO_SHAPEAPPMK)/connect-go/protoc-gen-connect-go.mk
