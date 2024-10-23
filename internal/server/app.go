package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/nakiner/go-service-template/internal/logger"
	swaggerui "github.com/nakiner/swagger-ui-go"
	"github.com/oklog/run"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type Config struct {
	HTTPPort                int           `envconfig:"SERVICE_PORT_HTTP" default:"8080"`
	GRPCPort                int           `envconfig:"SERVICE_PORT_GRPC" default:"8082"`
	HTTPDebugPort           int           `envconfig:"SERVICE_PORT_HTTP_DEBUG" default:"8084"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT" default:"15s"`
	GracefulShutdownDelay   time.Duration `envconfig:"GRACEFUL_SHUTDOWN_DELAY" default:"30s"`
	LogLevel                zapcore.Level `envconfig:"LOG_LEVEL" default:"info"`
}

type App struct {
	runGroup    run.Group
	http        *httpServer
	grpc        *grpcServer
	closer      *closer
	router      *chi.Mux
	debugRouter *chi.Mux
}

var appLogger = logger.Logger()

func fromEnv() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	fmt.Println(cfg)
	return cfg
}

func New() *App {
	loadLocalValuesYaml()
	appCfg := fromEnv()
	logger.SetLevel(appCfg.LogLevel)
	app := new(App)

	app.closer = new(closer)
	app.closer.add(func() error {
		return logger.Logger().Sync()
	})

	app.router = chi.NewRouter()
	app.debugRouter = chi.NewRouter()
	app.http = newHTTPServer(appCfg.HTTPPort, appCfg.HTTPDebugPort, appCfg.GracefulShutdownTimeout)
	app.http.server.Handler = app.router
	app.http.debugServer.Handler = app.debugRouter
	app.grpc = newGRPCServer(appCfg.GRPCPort, appCfg.GracefulShutdownTimeout)
	app.AddActor(interruptActor(appCfg.GracefulShutdownDelay))
	app.AddActor(app.http.actor())
	app.AddActor(app.http.debugActor())
	app.AddActor(app.closer.actor())

	return app
}

func (a *App) addGRPCServerActor() {
	a.AddActor(a.grpc.actor())
}

func (a *App) UseGrpcServerOptions(opt ...grpc.ServerOption) {
	a.grpc.opts = append(a.grpc.opts, opt...)
}

func (a *App) SetServeMux(mux *runtime.ServeMux) {
	a.router.Mount("/", mux)
	a.debugRouter.Mount("/", mux)
}

func (a *App) HTTP() *chi.Mux {
	return a.router
}

func (a *App) DebugHTTP() *chi.Mux {
	return a.debugRouter
}

func (a *App) WithSwaggerUI(swaggerJSON []byte) {
	a.debugRouter.Mount("/swagger", http.StripPrefix("/swagger", swaggerui.HTTPHandler()))
	a.debugRouter.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(swaggerJSON)
	})
	a.debugRouter.Handle("/swagger.json", handlerFunc)
}

func (a *App) WithProfiler() {
	a.debugRouter.Mount("/debug", middleware.Profiler())
}

func (a *App) Run() error {
	appLogger.Warn("application started")
	defer appLogger.Warn("application stopped")
	return a.runGroup.Run()
}

func (a *App) AddActor(execute func() error, interrupt func(error)) {
	a.runGroup.Add(execute, interrupt)
}

func (a *App) AddCloser(closer func() error) {
	a.closer.add(closer)
}

func (a *App) GRPC() *grpc.Server {
	if a.grpc.server == nil {
		a.grpc.setupServer()
		a.addGRPCServerActor()
	}

	return a.grpc.server
}
