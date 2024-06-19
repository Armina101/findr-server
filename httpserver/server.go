package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SrvConfig: struct to hold server configuration
type SrvConfig struct {
	Address        string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxHeaderBytes int
}

// Server: struct to hold server configuration
type Server struct {
	mux            *http.ServeMux
	middleware     []func(http.Handler) http.Handler
	logger         *slog.Logger
	shutdownSignal chan os.Signal
	srvConfig      SrvConfig
}

// ConfigServer: function to configure server
func ConfigServer(srvCfg *SrvConfig, lg *slog.Logger) *Server {
	return &Server{
		mux:            http.NewServeMux(),
		srvConfig:      *srvCfg,
		middleware:     []func(http.Handler) http.Handler{},
		logger:         lg,
		shutdownSignal: make(chan os.Signal),
	}
}
func (srv *Server) StartApp() {
	server := &http.Server{
		Addr:           srv.srvConfig.Address,
		Handler:        srv.mux,
		ReadTimeout:    srv.srvConfig.ReadTimeout,
		WriteTimeout:   srv.srvConfig.WriteTimeout,
		IdleTimeout:    srv.srvConfig.IdleTimeout,
		MaxHeaderBytes: srv.srvConfig.MaxHeaderBytes,

	}

	// Initializing a graceful shutdown of the application
	srv.ShutdownApp(server)
	// srv.logger.Info("Server started on port " + srv.srvConfig.Address)
	// srv.logger.Info(http.ListenAndServe(srv.srvConfig.Address, srv.mux))
}

func (srv *Server) ShutdownApp(server *http.Server) {

	// setting up context timeout to graceful shutdown of the application
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	signal.Notify(srv.shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {

		srv.logger.Info("Server started on port " + srv.srvConfig.Address)
		srv.logger.Info("Press Ctrl + C to stop the server using system calls")

		// Starting the server in a goroutine
		err := http.ListenAndServe(srv.srvConfig.Address, srv.mux)
		if err != nil && err != http.ErrServerClosed {
			srv.logger.Error("Server failed to start: ", err)
			os.Exit(1)
		}

	}()

	// Waiting for the shutdown signal (system calls)
	<-srv.shutdownSignal
	srv.logger.Info("Shutting down the Findr App server")
	if err := server.Shutdown(ctx); err != nil {
		srv.logger.Error("Server failed to shutdown Properly: ", err)
		os.Exit(1)
	}

	// Gracefully stopping the server
	srv.logger.Info("Server gracefully stopped properly")
	os.Exit(0)

}

// Use: function to add middlewares to the server easily
func (srv *Server) Use(middleware ...func(http.Handler) http.Handler) {
	srv.middleware = append(srv.middleware, middleware...)
}

func (srv *Server) AddMiddleware(handler http.Handler) http.Handler {
	for i := len(srv.middleware) - 1; i >= 0; i-- {
		handler = srv.middleware[i](handler)
	}
	return handler

}

// StaticFile: function to serve static files from the server to the client
func (srv *Server) StaticFile(path, file string) {
	srv.mux.Handle(path, http.StripPrefix("add the assest path", http.FileServer(http.Dir(file))))
}

// LoggingAppRequest: Basic custom middleware to track how long a request what processed
// since am not using no framework
func (srv *Server) LoggingAppRequest(lg *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
			begin := time.Now()
			lg.Info("%v  %v | Started .... ", rq.Method, rq.URL.Path)
			next.ServeHTTP(wr, rq)
			lg.Info("%s in %v | Completed ..... ", rq.URL.Path, time.Since(begin))

		})

	}
}
