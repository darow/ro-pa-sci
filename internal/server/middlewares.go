package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey uint8

const (
	ctxRequestIDKey = "request_id"
	ctxUserKey      = "user"
)

// setRequestID Присваиваем каждому запросу ID. Передаем его пользователю. Теперь мы можем находить каждый запрос по ID.
func (s *server) setRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Writer.Header().Set("X-Request-ID", id)
		c.Set(ctxRequestIDKey, id)
	}
}

func (s *server) auth(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session, err := s.store.Session().Find(token)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	user, err := s.store.User().Find(session.UserID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(ctxUserKey, user)
}
