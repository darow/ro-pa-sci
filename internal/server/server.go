package server

import (
	"github.com/darow/ro-pa-sci/internal/store"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/darow/ro-pa-sci/docs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	store  store.Store
	router *gin.Engine
	logger *zap.SugaredLogger
	hub    *wsHub
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
		filepath := "./frontend/distribute/" + c.Param("filename")
		c.File(filepath)
	})
    s.router.GET("/frontend/source/templates/:filename", func(c *gin.Context) {
		filepath := "./frontend/source/templates/" + c.Param("filename")
		c.File(filepath)
	})
    s.router.GET("/frontend/source/bootstrap/:filename", func(c *gin.Context) {
		filepath := "./frontend/source/bootstrap/" + c.Param("filename")
		c.File(filepath)
	})

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	s.router.POST("/user", s.createUser)
	s.router.POST("/session", s.createSession)
	s.router.GET("/online_users", s.getOnlineUsers)

	authorized := s.router.Group("/auth", s.auth)
	authorized.GET("/", s.whoAmI)
	authorized.GET("/logout", s.logout)
	authorized.GET("/ws", s.wsHandler())
}
