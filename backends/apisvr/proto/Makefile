.PHONY: setup
setup: buf.yaml

buf.yaml:
	buf mod init

gen: generate

.PHONY: generate
generate: buf.yaml
	buf generate

PROTOSET_BIN=protoset.bin
$(PROTOSET_BIN): build

.PHONY: protoset_path
protoset_path:
	@echo $(PROTOSET_BIN)

.PHONY: build
build: buf.yaml
	buf build -o $(PROTOSET_BIN)
