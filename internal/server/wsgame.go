package server

import "github.com/gorilla/websocket"

type wsGame struct {
	players map[int]*websocket.Conn
}
