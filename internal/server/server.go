package server

import (
	"net/http"

	"rock-paper-scissors/internal/store"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	store  store.Store
	router *gin.Engine
	logger *zap.SugaredLogger
}

// ServeHTTP server должен удовлетворять интерфейсу http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, logger *zap.SugaredLogger) *server {
	s := &server{
		store:  store,
		router: gin.Default(),
		logger: logger,
	}

	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID())

	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/game.html") })
	s.router.GET("/:filename", func(c *gin.Context) {
		filepath := "./web/" + c.Param("filename")
		c.File(filepath)
	})

	s.router.POST("/user", s.createUser)
	s.router.POST("/session", s.createSession)
	s.router.GET("/online_users", s.getOnlineUsers)

	authorized := s.router.Group("/auth", s.auth)
	authorized.GET("/", s.whoAmI)
	authorized.GET("/logout", s.logout)
	authorized.GET("/ws", s.wsHandler())
}
