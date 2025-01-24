package server

import (
	"github.com/osamikoyo/router/internal/parser"
	"github.com/osamikoyo/router/pkg/loger"
	"net/http"
)

type Handler struct {
	Logger loger.Logger
	Config parser.Config
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h *Handler) getError(handler HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := handler(writer, request);err != nil{
			h.Logger.Error().Err(err)
		}
	}
}

func (s Server) Handler(w http.ResponseWriter, r *http.Request) error {
	
}