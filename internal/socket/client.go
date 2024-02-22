/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-22 00:12
 * @Description:
 */

package socket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leafney/flutter-assets-helper/pkg/protocol"
	"github.com/leafney/flutter-assets-helper/pkg/vars"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()

		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error(err)
			c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)
		// 判断消息类型
		if msg.Type == vars.HeartBeat {
			pong := &protocol.Message{
				Content: "pong",
				Type:    vars.HeartBeat,
			}
			pontByte, err := proto.Marshal(pong)
			if err != nil {
				log.Error(err)
			}
			if err := c.Conn.WriteMessage(websocket.BinaryMessage, pontByte); err != nil {
				log.Error(err)
			}
		} else {
			//	普通消息
			XServer.Operation <- message
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
			log.Error(err)
		}
	}
}
