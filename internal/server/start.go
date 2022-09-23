package server

import (
	"net"
	"net/http"

	"go.uber.org/zap"
)

func Start(cfg *Config, logger *zap.SugaredLogger) error {
	s := newServer(logger)
	//http.HandleFunc("/ws", wsEndpoint)

	return http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), s)
}
