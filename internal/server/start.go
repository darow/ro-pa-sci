package server

import (
	"github.com/darow/ro-pa-sci/internal/store/teststore"
	"net"
	"net/http"

	"go.uber.org/zap"
)

func Start(cfg *Config, logger *zap.SugaredLogger) error {
	store := teststore.New()

	s := newServer(store, logger)

	return http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), s)
}
