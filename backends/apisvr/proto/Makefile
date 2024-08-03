.PHONY: default
default: build lint test

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

include $(PATH_TO_SHAPEAPPMK)/buf/buf.mk
include $(PATH_TO_SHAPEAPPMK)/buf/proto/base.mk
include $(PATH_TO_SHAPEAPPMK)/protobuf/protoc-gen-go.mk
include $(PATH_TO_SHAPEAPPMK)/connect-go/protoc-gen-connect-go.mk

.PHONY: build
build: buf-build

.PHONY: lint
lint: buf-lint

.PHONY: generate
generate: buf-generate

include $(PATH_TO_SHAPEAPPMK)/grpcurl/base.mk

.PHONY: test-connections
test-connections: $(GRPCURL_CLI) $(BUF_CLI) build
	./test-connections.sh
