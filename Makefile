
LOCAL_BIN := $(CURDIR)/bin

BUF_BUILD := $(LOCAL_BIN)/buf
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.bin-generate:
	$(info Generating protobuf files...)
	PATH="$(LOCAL_BIN):$(PATH)" $(BUF_BUILD) generate

generate: .bin-generate .tidy

.tidy:
	go mod tidy

.PHONY: \
	bin-deps

