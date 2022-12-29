package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept multipart/form-data
// @Produce json
// @Param input formData model.User true "login-pass"
// @Success 201 {object} model.User
// @Success 400 {object} any "user data"
// @Router /user [post]
func (s *server) createUser(c *gin.Context) {
	u := &model.User{
		Username: c.PostForm("name"),
		Password: c.PostForm("password"),
	}

	err := s.store.User().Create(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	s.createSession(c)
}

// @Summary Login
// @Tags auth
// @Description create session for existing account
// @ID create-session
// @Accept multipart/form-data
// @Produce json
// @Param input formData model.User true "login-pass"
// @Success 201 {object} any "user data"
// @Success 400 {object} any
// @Router /session [post]
func (s *server) createSession(c *gin.Context) {
	u, err := s.store.User().Login(c.PostForm("name"), c.PostForm("password"))
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

// @Summary WhoAmI
// @Tags auth
// @Description check current user
// @ID who-am-i
// @Success 200 {object} any "user data"
// @Success 400 {object} any
// @Router /auth/ [get]
func (s *server) whoAmI(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary GetInvites
// @Tags auth
// @Description get invites that sent from me and to me
// @ID get-invites
// @Success 200 {object} any "invites"
// @Success 400 {object} any
// @Router /invites [get]
func (s *server) getInvites(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		return
	}

	invites, err := s.store.Invite().GetByUser(user.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, invites)
}

// @Summary Logout
// @Tags auth
// @Description delete session cookie
// @ID logout
// @Success 200
// @Success 400 {object} any
// @Router /auth/logout [get]
func (s *server) logout(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		return
	}

	user.IsOnline = false
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(http.StatusOK)
}
