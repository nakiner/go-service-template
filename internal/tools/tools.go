//go:build tools
// +build tools

package tools

// list packages here to prevent them from removal out of go.mod
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	_ "github.com/googleapis/googleapis/google/example/endpointsapis/goapp"
	_ "github.com/kannman/modtools"

	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"

	_ "github.com/bufbuild/buf/cmd/buf"
)
