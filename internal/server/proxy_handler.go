package server

import (
	"context"
	"fmt"
	"github.com/osamikoyo/router/internal/parser"
	"github.com/osamikoyo/router/pkg/loger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func (h Handler) ProxyHandler(w http.ResponseWriter, r *http.Request) error {
	domain := strings.Split(r.Host, ":")[0]

	port, ok := h.Config.Lines[domain]
	if !ok {
		return nil
	}

	addr := fmt.Sprintf("%s:%d", "localhost", port)
	ctx,cancel := context.WithTimeout(context.Background(), 50 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, r.Method, addr, r.Body)

	req.URL.Scheme = r.URL.Scheme

	if err != nil{
		return err
	}

	client := &http.Client{
		Timeout: 50 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil{
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	_, err = w.Write(body)
	return err
}