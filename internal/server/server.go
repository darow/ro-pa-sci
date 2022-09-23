package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	logger *zap.SugaredLogger
	router *gin.Engine
}

// ServeHTTP server должен удовлетворять интерфейсу http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(logger *zap.SugaredLogger) *server {
	s := &server{
		logger: logger,
		router: gin.Default(),
	}

	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.Use()
	s.router.GET("/ws", s.wsHandler())
	s.router.GET("/:filename", func(c *gin.Context) {
		filepath := "./web/" + c.Param("filename")
		c.File(filepath)
	})
}
