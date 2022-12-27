package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/darow/ro-pa-sci/internal/model"

	"github.com/gorilla/websocket"
)

type wsHub struct {
	sync.Mutex
	userCons map[int]*websocket.Conn
}

func (h *wsHub) AddUser(userID int, con *websocket.Conn) {
	h.Lock()
	h.userCons[userID] = con
	h.Unlock()
}

func (h *wsHub) PopUser(userID int) {
	h.Lock()
	delete(h.userCons, userID)
	h.Unlock()
}

type wsRequest struct {
	messageType int
	userFrom    *model.User
	Action      string          `json:"action"`
	Body        json.RawMessage `json:"body"`
}

type wsResponse struct {
	Code int    `json:"code,omitempty"`
	Body string `json:"body,omitempty"`
}

func (s *server) handleWS(user *model.User, conn *websocket.Conn) {
	s.hub.AddUser(user.ID, conn)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			user.IsOnline = false
			s.hub.PopUser(user.ID)
			s.logger.Error(err)
			return
		}
		fmt.Println(string(p))

		request := wsRequest{
			messageType: messageType,
			userFrom:    user,
		}
		err = json.Unmarshal(p, &request)

		resp := s.wsProcess(request)

		buf, err := json.Marshal(resp)
		if err != nil {
			if err = conn.WriteMessage(messageType, []byte("error:"+string(err.Error()))); err != nil {
				s.logger.Error(err)
				return
			}
		}

		if err = conn.WriteMessage(messageType, buf); err != nil {
			s.logger.Error(err)
			return
		}
	}
}

func (s *server) wsProcess(r wsRequest) *wsResponse {
	switch r.Action {
	case "invite":
		return s.createInvite(r)
	}

	return &wsResponse{Code: http.StatusBadRequest, Body: "action not found"}
}

func (s *server) createInvite(r wsRequest) *wsResponse {
	var invite model.Invite
	err := json.Unmarshal(r.Body, &invite)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	invite.From = r.userFrom.ID
	invite.Timestamp = time.Now()

	conTo, ok := s.hub.userCons[invite.To]
	if !ok {
		return &wsResponse{Code: http.StatusBadRequest, Body: "Игрока с таким id сейчас нет в сети"}
	}

	// Обработали wsRequest. Изменяем его для отправки
	body := struct {
		From int `json:"from"`
	}{
		invite.From,
	}
	r.Body, err = json.Marshal(body)
	if err != nil {
		return &wsResponse{Code: http.StatusInternalServerError, Body: err.Error()}
	}

	payload, err := json.Marshal(r)
	if err != nil {
		return &wsResponse{Code: http.StatusInternalServerError, Body: err.Error()}
	}

	if err = conTo.WriteMessage(r.messageType, payload); err != nil {
		s.logger.Warn(err)
		return &wsResponse{Code: http.StatusInternalServerError, Body: err.Error()}
	}

	return &wsResponse{Code: http.StatusOK, Body: "пригласили игрока № " + strconv.Itoa(r.userFrom.ID)}
}
