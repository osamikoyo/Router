package server

import (
	"fmt"
	"github.com/osamikoyo/router/internal/parser"
	"github.com/osamikoyo/router/pkg/loger"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
	Logger loger.Logger
	Handlers Handler
}

func New() (Server, error) {
	cfg, err := parser.New().Parse()
	return Server{
		HttpServer: &http.Server{
			Addr: fmt.Sprintf("localhost:%d", cfg.Port),
		},
		Logger: loger.New(),
		Handlers: Handler{
			Logger: loger.New(),
			Config: cfg,
		},
	}, err
}

func (s Server) Run() error {
	s.Logger.Info().Msg("Starting routing server...")
	
	mux := http.NewServeMux()
	mux.Handle("/", s.Handlers.getError(s.Handlers.ProxyHandler))
	s.HttpServer.Handler = mux

	s.Logger.Info().Msg("Route http server started on " + fmt.Sprintf("%s:%d", "localhost", s.Handlers.Config.Port))

	return s.HttpServer.ListenAndServe()
}