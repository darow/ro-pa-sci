package server

import (
	"fmt"
	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetUsersList
// @Tags playersList
// @Description get list of all users
// @ID users-list
// @Produce json
// @Success 200 {object} any "user data"
// @Success 400 {object} any
// @Router /online_users [get]
func (s *server) getOnlineUsers(c *gin.Context) {
	c.JSON(http.StatusOK, s.store.User().GetTop())
}

// wsHandler ... TODO сделать документацию через свагер
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
		s.handleWS(user, ws)
	}
}
