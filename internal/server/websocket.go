package server

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

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
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
		}

		s.logger.Infoln("client connected")
		err = ws.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			log.Println(err)
		}

		reader(ws)
	}
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte("message recieved")); err != nil {
			log.Println(err)
			return
		}

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if err := conn.WriteMessage(messageType, []byte(text)); err != nil {
			log.Println(err)
			return
		}

	}
}
