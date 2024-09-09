LOCAL_BIN=$(CURDIR)/bin

RUN_ARGS:=
ifneq (,$(wildcard env/local.yaml))
    RUN_ARGS=--config=env/local.yaml
endif

PROTOC_BIN=protoc

PROTOC_GEN_GO_BIN=$(LOCAL_BIN)/protoc-gen-go
$(PROTOC_GEN_GO_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/golang/protobuf/protoc-gen-go


PROTOC_GEN_GO_GRPC_BIN=$(LOCAL_BIN)/protoc-gen-go-grpc
$(PROTOC_GEN_GO_GRPC_BIN):
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


PROTOC_GEN_GRPC_GATEWAY_BIN=$(LOCAL_BIN)/v2/protoc-gen-grpc-gateway
$(PROTOC_GEN_GRPC_GATEWAY_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway


PROTOC_GEN_OPENAPI_BIN=$(LOCAL_BIN)/v2/protoc-gen-openapiv2
$(PROTOC_GEN_OPENAPI_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2


BUF_BIN=$(LOCAL_BIN)/buf
$(BUF_BIN): $(PROTOC_GEN_BUF_BREAKING_BIN) $(PROTOC_GEN_BUF_LINT_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf

PROTOC_GEN_BUF_BREAKING_BIN=$(LOCAL_BIN)/protoc-gen-buf-breaking
$(PROTOC_GEN_BUF_BREAKING_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking

PROTOC_GEN_BUF_LINT_BIN=$(LOCAL_BIN)/protoc-gen-buf-lint
$(PROTOC_GEN_BUF_LINT_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint


MODTOOLS_BIN=$(LOCAL_BIN)/modtools
$(MODTOOLS_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/kannman/modtools


.protodeps: $(PROTOC_GEN_GO_BIN) $(PROTOC_GEN_GO_GRPC_BIN) $(PROTOC_GEN_GRPC_GATEWAY_BIN) $(PROTOC_GEN_OPENAPI_BIN) $(MODTOOLS_BIN) $(WRAPMUX_BIN) $(BUF_BIN)
	$(info all proto deps installed)

.PHONY: .vendorpb
.vendorpb:
	$(GOENV) go mod download
	rm -rf vendor.pb && $(MODTOOLS_BIN) vendor '**/*.proto' && mv vendor vendor.pb

.PHONY: generate
generate: .protodeps .vendorpb
	PATH=$(LOCAL_BIN):$(PATH) $(BUF_BIN) generate -v --path=api/go_service_template

.PHONY: run
run:
	go run cmd/service/main.go $(RUN_ARGS)

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -o bin/service ./cmd/service

.PHONY: modcache-clean
modcache-clean:
	go clean -cache -modcache
