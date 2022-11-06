package server

import (
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
