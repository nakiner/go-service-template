package main

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	hndl "github.com/nakiner/go-service-template/internal/handler/go_service_template/v1"
	pb "github.com/nakiner/go-service-template/pkg/pb/go_service_template/v1"
	"google.golang.org/grpc"

	"github.com/nakiner/go-service-template/internal/logger"
	"github.com/nakiner/go-service-template/internal/server"
)

func main() {
	app := server.New()
	initApp(app)
	mustInit(app.Run())
}

func initApp(app *server.App) {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	handler := hndl.NewService()

	app.SetHandler(mux)
	app.Use(middleware.Recoverer)
	app.Use(logger.RequestLogger())
	app.UseGrpcServerOptions(
		grpc.ChainUnaryInterceptor(
			server.WithUnaryServerRecovery(),
			logger.UnaryServerInterceptorLogger(),
		),
		grpc.ChainStreamInterceptor(
			server.WithStreamServerRecovery(),
			logger.StreamServerInterceptorLogger(),
		),
	)

	mustInit(pb.RegisterGoServiceTemplateServiceV1HandlerServer(ctx, mux, handler))
	pb.RegisterGoServiceTemplateServiceV1Server(app.GRPC(), handler)
}

func mustInit(err error) {
	if err != nil {
		logger.FatalKV(context.Background(), "init failure", "err", err)
	}
}
