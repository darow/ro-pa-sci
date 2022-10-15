package server

import (
	"net"
	"net/http"

	"rock-paper-scissors/internal/store/teststore"

	"go.uber.org/zap"
)

func Start(cfg *Config, logger *zap.SugaredLogger) error {
	store := teststore.New()

	s := newServer(store, logger)

	return http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), s)
}
