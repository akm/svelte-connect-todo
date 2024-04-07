.PHONY: setup
setup: buf.yaml

buf.yaml:
	buf mod init

gen: generate

.PHONY: generate
generate: buf.yaml
	buf generate
