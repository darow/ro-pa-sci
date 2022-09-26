package server

import (
	"net/http"

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
			c.AbortWithError(http.StatusBadRequest, err)
		}

		u.Sanitize()
		c.JSON(http.StatusCreated, u)
	}
}
