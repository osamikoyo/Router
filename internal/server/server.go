package server

import (
	"github.com/osamikoyo/router/pkg/loger"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
	Logger loger.Logger
	Handlers Handler
}