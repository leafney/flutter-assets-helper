/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-05 22:48
 * @Description:
 */

package run

import (
	"bufio"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/leafney/flutter-assets-helper/internal/socket"
	"github.com/leafney/flutter-assets-helper/web"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the Fah web project",
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		port, _ := cmd.Flags().GetInt("port")

		if daemon {
			// Start Fiber web project in background (persistent execution)
			go StartWeb(port)
		} else {
			// Start Fiber web project and open a web page
			StartWeb(port)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		log.Printf("file %v", file)
		loadPubspec3(file)
	},
}

func init() {
	runCmd.Flags().BoolP("daemon", "d", false, "Run in daemon mode (persistent execution)")
	runCmd.Flags().IntP("port", "p", 8080, "Port number to run the Gin web project")
	runCmd.Flags().StringP("file", "f", "pubspec.yaml", "Flutter project default pubspec.yaml file")
	rootCmd.AddCommand(runCmd)
}

func StartWeb(port int) {
	app := fiber.New()

	log.Println("app start")

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello World")
	//})

	// websocket
	app.Use("/ws", func(c *fiber.Ctx) error {
		log.Println("调用了 websocket")

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	//wsConf := websocket.Config{
	//	HandshakeTimeout: 100 * time.Second,
	//	Origins: []string{
	//		"http://localhost:8080",
	//		"http://127.0.0.1:8080",
	//	},
	//}

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		client := &socket.Client{
			Conn: c,
			Send: make(chan []byte),
		}
		// 处理消息的接收和发送
		go client.Read()
		go client.Write()

		// 这里要一直阻塞
		select {}
	}))

	//app.Get("/ws", websocket.New(func(c *websocket.Conn) {
	//	// Access the *websocket.Conn methods
	//	// For example:
	//	// c.Locals("allowed") // true
	//	// c.Params("id") // 123
	//	// c.Query("v") // 1.0
	//	// c.Cookies("session") // ""
	//	// c.ReadMessage()
	//	// c.WriteMessage()
	//
	//	//log.Println(c.Params("id"))
	//	log.Println("websocket 收到了消息")
	//
	//	var (
	//		mt  int
	//		msg []byte
	//		err error
	//	)
	//	for {
	//		if mt, msg, err = c.ReadMessage(); err != nil {
	//			log.Println("read:", err)
	//			break
	//		}
	//		log.Printf("文件类型为 %v", mt)
	//
	//		//log.Printf("recv: %s", msg)
	//
	//
	//		/*
	//			// 直接上传图片文件后，接收并保存到本地
	//			reader := bytes.NewReader(msg)
	//			bts, err := io.ReadAll(reader)
	//			err = os.WriteFile("abc.png", bts, 0644)
	//			if err != nil {
	//				log.Fatalln(err)
	//			}
	//		*/
	//
	//		message := &protocol.Message{}
	//		proto.Unmarshal(msg, message)
	//
	//		//log.Println(rose.JsonMarshalStr(message))
	//
	//		bts, err := io.ReadAll(bytes.NewReader(message.File))
	//		err = os.WriteFile("abcd.png", bts, 0644)
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//
	//		//if err = c.WriteMessage(mt, msg); err != nil {
	//		//	log.Println("write:", err)
	//		//	break
	//		//}
	//	}
	//}, wsConf))

	// webui
	uiDist, err := fs.Sub(web.UiStatic, "dist")
	if err != nil {

	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(uiDist),
	}))

	go func() {
		go socket.XServer.Start()

		newPort, canUsed := getAvailablePort(port)
		if !canUsed {
			fmt.Printf("[Warn] Port number %d is occupied, automatically match port number: %d \n", port, newPort)
		}

		ip := getLocalIP()
		fmt.Printf("App running on:\n\n→  Local:    http://127.0.0.1:%d \n→  NetWork:  http://%s:%d\n\n", newPort, ip, newPort)

		addr := fmt.Sprintf(":%d", newPort)
		if err := app.Listen(addr); err != nil {
			log.Fatal("Failed to start Fiber web project: ", err)
		}
	}()

	select {}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, address := range addrs {
		// 检查IP地址是否为IPv4地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}

func getAvailablePort(port int) (int, bool) {
	// 标识提供的端口号是否可用
	canUse := true
	for {
		conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)))
		if err != nil {
			return port, canUse
		}
		_ = conn.Close()
		port++
		canUse = false
	}
}

func loadPubspec3(filePath string) {
	fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		fmt.Println(line)
		if line == "" {
			fmt.Println("空白行")
		}

	}
}
