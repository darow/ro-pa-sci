package server

import (
	"fmt"
	"github.com/darow/ro-pa-sci/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	Username string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body createUserInput true "account-info"
// @Success 201 {bool} bool true
// @Success 400 {object} map[string]string
// @Failure default {object} map[string]string
// @Router /user [post]
func (s *server) createUser(c *gin.Context) {
	var input createUserInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	u := &model.User{
		Username: input.Username,
		Password: input.Password,
	}

	err := s.store.User().Create(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, true)
}

func (s *server) createSession(c *gin.Context) {
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
	c.JSON(http.StatusCreated, u)
}

func (s *server) whoAmI(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
	}

	c.JSON(http.StatusCreated, user)
}

func (s *server) getOnlineUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, s.store.User().GetTop())
}

func (s *server) logout(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
	}

	user.IsOnline = false
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(http.StatusOK)
}
