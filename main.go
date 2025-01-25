package main

import (
	server2 "github.com/osamikoyo/router/internal/server"
	"github.com/osamikoyo/router/pkg/loger"
)

func main(){
	logger := loger.New()
	server, err := server2.New()
	if err != nil{
		logger.Error().Err(err)
	}
	logger.Info().Msg("starting...")
	if err := server.Run();err != nil{
		logger.Error().Err(err)
	}
}