package server

import (
	"net/http"
	"time"

	"rock-paper-scissors/internal/model"

	"github.com/gin-gonic/gin"
)

func (s *server) handleCreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := &model.User{
			Login:    c.PostForm("login"),
			Password: c.PostForm("password"),
		}

		err := s.store.User().Create(u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, true)
	}
}

func (s *server) handleCreateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := s.store.User().Login(c.PostForm("login"), c.PostForm("password"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		session, err := s.store.Session().Create(u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		c.SetCookie("session", session.Token, int(session.ExpirationTime.Sub(time.Now()).Seconds()), "", "", false, false)
		c.JSON(http.StatusCreated, true)
	}
}
