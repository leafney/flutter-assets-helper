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
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello World")
	//})

	// webui
	uiDist, err := fs.Sub(web.UiStatic, "dist")
	if err != nil {

	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(uiDist),
	}))

	go func() {
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
