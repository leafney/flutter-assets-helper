/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-22 11:45
 * @Description:
 */

package socket

import (
	"github.com/gofiber/contrib/websocket"
	"log"
)

var XServer = NewServer()

type Server struct {
	Client     *Client
	Operation  chan []byte
	Register   chan *websocket.Conn
	UnRegister chan *websocket.Conn
}

func NewServer() *Server {
	return &Server{
		Client:     nil,
		Operation:  make(chan []byte),
		Register:   make(chan *websocket.Conn),
		UnRegister: make(chan *websocket.Conn),
	}
}

func (s *Server) Start() {
	for {
		select {

		case message := <-s.Operation:
			log.Printf("接收到了 %v", string(message))

			//msg := &protocol.Message{}
			//proto.Unmarshal(message, msg)
			//
			//log.Println(rose.JsonMarshalStr(msg))
		}
	}
}
