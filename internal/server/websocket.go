package server

import (
	"fmt"
	"github.com/darow/ro-pa-sci/internal/model"
	"sync"

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

		if err := conn.WriteMessage(messageType, []byte("message recieved")); err != nil {
			s.logger.Error(err)
			return
		}
	}
}
