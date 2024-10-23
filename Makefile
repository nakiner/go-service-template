LOCAL_BIN=$(CURDIR)/bin

RUN_ARGS:=
ifneq (,$(wildcard env/local.yaml))
    RUN_ARGS=--config=env/local.yaml
endif

PROTOC_GEN_GO_BIN=$(LOCAL_BIN)/protoc-gen-go
$(PROTOC_GEN_GO_BIN):
	[ -f $(PROTOC_GEN_GO_BIN) ] || GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go


PROTOC_GEN_GO_GRPC_BIN=$(LOCAL_BIN)/protoc-gen-go-grpc
$(PROTOC_GEN_GO_GRPC_BIN):
	[ -f $(PROTOC_GEN_GO_GRPC_BIN) ] || GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


PROTOC_GEN_GRPC_GATEWAY_BIN=$(LOCAL_BIN)/v2/protoc-gen-grpc-gateway
$(PROTOC_GEN_GRPC_GATEWAY_BIN):
	[ -f $(PROTOC_GEN_GRPC_GATEWAY_BIN) ] || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway


PROTOC_GEN_OPENAPI_BIN=$(LOCAL_BIN)/v2/protoc-gen-openapiv2
$(PROTOC_GEN_OPENAPI_BIN):
	[ -f $(PROTOC_GEN_OPENAPI_BIN) ] || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2


BUF_BIN=$(LOCAL_BIN)/buf
BUF_BIN_TAG=v1.40.1
BUF_BIN_URL=https://github.com/bufbuild/buf/releases/download/$(BUF_BIN_TAG)/buf-$(shell uname -s)-$(shell uname -m).tar.gz
$(BUF_BIN): $(LOCAL_BIN)
	[ -f $(BUF_BIN) ] || curl -sSL $(BUF_BIN_URL) | tar -C $(LOCAL_BIN) --strip-components 2 -xz buf/bin/buf


MODTOOLS_BIN=$(LOCAL_BIN)/modtools
$(MODTOOLS_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/kannman/modtools

GOBINDATA_BIN=$(LOCAL_BIN)/go-bindata
$(GOBINDATA_BIN):
	[ -f $(GOBINDATA_BIN) ] || GOBIN=$(LOCAL_BIN) go install github.com/kevinburke/go-bindata/v4/...@latest


.protodeps: $(PROTOC_GEN_GO_BIN) $(PROTOC_GEN_GO_GRPC_BIN) $(PROTOC_GEN_GRPC_GATEWAY_BIN) $(PROTOC_GEN_OPENAPI_BIN) $(MODTOOLS_BIN) $(WRAPMUX_BIN) $(BUF_BIN) $(GOBINDATA_BIN)
	$(info all proto deps installed)

.PHONY: .vendorpb
.vendorpb:
	$(GOENV) go mod download
	rm -rf vendor.pb && $(MODTOOLS_BIN) vendor '**/*.proto' && mv vendor vendor.pb

.PHONY: generate
generate: .protodeps .vendorpb
	PATH=$(LOCAL_BIN):$(PATH) $(BUF_BIN) generate --path=api/go_service_template
	rm -rf vendor.pb
	$(GOBINDATA_BIN) -pkg bindata -o internal/pkg/bindata/swagger-json.go api/api.swagger.json

.PHONY: run
run:
	go run cmd/service/main.go $(RUN_ARGS)

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -o bin/service ./cmd/service

.PHONY: modcache-clean
modcache-clean:
	go clean -cache -modcache
