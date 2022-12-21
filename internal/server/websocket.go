package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/darow/ro-pa-sci/internal/model"

	"github.com/gorilla/websocket"
)

type wsHub struct {
	sync.Mutex
	users map[int]*websocket.Conn
}

func (h *wsHub) AddUser(userID int, con *websocket.Conn) {
	h.Lock()
	h.users[userID] = con
	h.Unlock()
}

func (h *wsHub) PopUser(userID int) {
	h.Lock()
	delete(h.users, userID)
	h.Unlock()
}

type wsRequest struct {
	userID int
	Action string `json:"action,omitempty"`
	Body   string `json:"body,omitempty"`
}

type wsResponse struct {
	Code int    `json:"code"`
	Body string `json:"body"`
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

		var request wsRequest
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
	case "sendGameInvite":
		s.sendGameInvite(r)
	}

	return &wsResponse{Code: http.StatusBadRequest, Body: "action not found"}
}

func (s *server) sendGameInvite(r wsRequest) *wsResponse {
	return &wsResponse{Code: http.StatusOK, Body: "Траляля пригласили игрока с id(нет)" + strconv.Itoa(r.userID)}
}
