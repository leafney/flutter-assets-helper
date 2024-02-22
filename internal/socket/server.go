/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-22 11:45
 * @Description:
 */

package socket

import (
	"bytes"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/leafney/flutter-assets-helper/pkg/protocol"
	"github.com/leafney/flutter-assets-helper/pkg/vars"
	"github.com/leafney/rose"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
)

var XServer = NewServer()

type Server struct {
	Client    *Client
	Operation chan []byte
	Register  chan *websocket.Conn
	Ungister  chan *websocket.Conn
}

func NewServer() *Server {
	return &Server{
		Client:    nil,
		Operation: make(chan []byte),
		Register:  make(chan *websocket.Conn),
		Ungister:  make(chan *websocket.Conn),
	}
}

func (s *Server) Start() {
	for {
		select {
		//case conn := <-s.Register:
		//case conn:=<-s.Ungister:

		case message := <-s.Operation:
			//log.Printf("接收到了 %v", string(message))

			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			//log.Println(rose.JsonMarshalStr(msg))

			if msg.ContentType == vars.ContentTypeText {
				value := msg.GetContent()
				log.Printf("获取到内容 %v", value)
			} else if msg.ContentType == vars.ContentTypeColor {
				color := msg.GetContent()
				log.Printf("获取到色值 %v", color)

			} else if msg.ContentType == vars.ContentTypeImage {
				// TODO 图片和文件上传都是走的一个逻辑，判断是图片还是压缩文件，需要在后端去完成

				fileName := msg.FileName
				if rose.StrIsEmpty(fileName) {
					fileName = fmt.Sprintf("%v.png", rose.RandStr(10))
				}

				bts, err := io.ReadAll(bytes.NewReader(msg.File))

				os.MkdirAll("./tmp", os.ModePerm)
				err = os.WriteFile(rose.FJoinPath("./tmp", fileName), bts, 0644)
				if err != nil {
					log.Fatalln(err)
				}
				log.Printf("文件 %v 保存成功", fileName)

			} else if msg.ContentType == vars.ContentTypeZipFile {

			} else {

			}
		}
	}
}
