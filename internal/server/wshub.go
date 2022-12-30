package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/darow/ro-pa-sci/internal/model"
	"go.uber.org/zap"

	"github.com/gorilla/websocket"
)

type wsHub struct {
	sync.Mutex
	userCons map[int]*websocket.Conn
	games    map[int]*wsGame
	logger   *zap.SugaredLogger
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

func (h *wsHub) writeResponseWS(conn *websocket.Conn, resp *wsResponse) {
	buf, err := json.Marshal(resp)
	if err != nil {
		err = fmt.Errorf("ошибка при записи ответа. %w", err)
		if err = conn.WriteMessage(1, []byte("error:"+string(err.Error()))); err != nil {
			h.logger.Error(err)
			return
		}
	}

	if err = conn.WriteMessage(1, buf); err != nil {
		h.logger.Error(err)
		return
	}
}

func (h *wsHub) broadcast(msg string, cons ...*websocket.Conn) {
	for _, con := range cons {
		con.WriteMessage(1, []byte(msg))
	}
}

func (h *wsHub) StartGame(con1, con2 *websocket.Conn) (model.Game, error) {
	game := wsGame{}
	_ = game

	h.broadcast("game starts. send me \"rdy\"")

	return model.Game{}, nil
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

		s.hub.writeResponseWS(conn, resp)
	}
}

func (s *server) wsProcess(r wsRequest) *wsResponse {
	switch r.Action {
	case "createInvite":
		return s.createInvite(r)
	case "decideInvite":
		return s.decideInvite(r)
	}

	return &wsResponse{Code: http.StatusBadRequest, Body: "action not found"}
}

func (s *server) decideInvite(r wsRequest) *wsResponse {
	var request struct {
		InviteID int   `json:"inviteID"`
		Decision uint8 `json:"decision"`
	}
	err := json.Unmarshal(r.Body, &request)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	if _, ok := model.Decisions[request.Decision]; !ok {
		return &wsResponse{Code: http.StatusBadRequest, Body: "нет такого ответа на приглашение"}
	}

	inv, err := s.store.Invite().Get(request.InviteID)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	if inv.Decision != model.DecisionNotDecided {
		return &wsResponse{Code: http.StatusBadRequest, Body: "вы уже ответили на это приглашение"}
	}

	if inv.To != r.userFrom.ID {
		return &wsResponse{Code: http.StatusForbidden, Body: fmt.Errorf("%w вы не являетесь получателем приглашения", ErrForbidden).Error()}
	}

	if request.Decision == model.DecisionAccepted {
		var ok bool
		var con1, con2 *websocket.Conn
		if con1, ok = s.hub.userCons[r.userFrom.ID]; !ok {
			return &wsResponse{Code: http.StatusBadRequest, Body: "на сервере нет вашего подключения. обновите страницу"}
		}

		if con2, ok = s.hub.userCons[inv.From]; !ok {
			return &wsResponse{Code: http.StatusBadRequest, Body: "пользователь отправивший приглашение не онлайн"}
		}

		s.hub.writeResponseWS(con1, &wsResponse{Code: http.StatusOK})

		var game model.Game
		// возвращать ошибку из startGame в самом крайнем случае, чтобы записать decision
		game, err = s.hub.StartGame(con1, con2)
		if err != nil {
			return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
		}

		err = s.store.Game().Create(&game)
		if err != nil {
			return &wsResponse{Code: http.StatusBadRequest, Body: fmt.Errorf("результат игры не был записан %w", err).Error()}
		}
	}

	inv.Decision = request.Decision
	err = s.store.Invite().Update(inv)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	return &wsResponse{Code: http.StatusOK}
}

func (s *server) createInvite(r wsRequest) *wsResponse {
	var invite model.Invite
	err := json.Unmarshal(r.Body, &invite)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	if invite.To == r.userFrom.ID {
		return &wsResponse{Code: http.StatusBadRequest, Body: "from == to. нельзя пригласить себя"}
	}

	invite.From = r.userFrom.ID

	conTo, ok := s.hub.userCons[invite.To]
	if !ok {
		return &wsResponse{Code: http.StatusBadRequest, Body: "игрока с таким id сейчас нет в сети"}
	}

	err = s.store.Invite().Create(&invite)
	if err != nil {
		return &wsResponse{Code: http.StatusBadRequest, Body: err.Error()}
	}

	r.Body, err = json.Marshal(invite)
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
