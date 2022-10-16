package server

import (
	"fmt"
	"github.com/darow/ro-pa-sci/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (s *server) wsHandler() func(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return func(c *gin.Context) {
		u, ok := c.Get(ctxUserKey)
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		}
		user, ok := u.(*model.User)
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		}

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user.IsOnline = true
		s.reader(ws, user)
	}
}

func (s *server) reader(conn *websocket.Conn, user *model.User) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			user.IsOnline = false
			s.logger.Error(err)
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte("message recieved")); err != nil {
			s.logger.Error(err)
			return
		}
	}
}
