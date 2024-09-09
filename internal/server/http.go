package server

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type httpServer struct {
	listener        net.Listener
	server          *http.Server
	port            int
	gracefulTimeout time.Duration
}

func newHTTPServer(port int, gracefulTimeout time.Duration) *httpServer {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(errors.Wrap(err, "http listener init failure"))
	}
	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
	}

	r := &httpServer{
		listener:        listener,
		server:          server,
		port:            port,
		gracefulTimeout: gracefulTimeout,
	}

	return r
}

func (r *httpServer) actor() (func() error, func(error)) {
	return func() error {
			appLogger.Warnw("http server started", "port", r.port)
			return errors.Wrap(r.server.Serve(r.listener), "http server")
		}, func(err error) {
			ctx, cancel := context.WithTimeout(context.Background(), r.gracefulTimeout)
			defer cancel()

			r.server.SetKeepAlivesEnabled(false)
			if err := errors.Wrap(r.server.Shutdown(ctx), "http server: error during shutdown"); err != nil {
				appLogger.Error("http server: stop failure", err)
			} else {
				appLogger.Warn("http server: gracefully stopped")
			}
		}
}
