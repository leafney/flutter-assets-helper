/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-22 11:45
 * @Description:
 */

package socket

import (
	"log"
)

var XServer = NewServer()

type Server struct {
	Client    *Client
	Operation chan []byte
}

func NewServer() *Server {
	return &Server{
		Client:    nil,
		Operation: make(chan []byte),
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
