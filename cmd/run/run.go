/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-05 22:48
 * @Description:
 */

package run

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
	"net"
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
}

func StartWeb(port int) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	go func() {
		ip := getLocalIP()
		fmt.Printf("App running on:\n\n→  Local:    http://127.0.0.1:%d \n→  NetWork:  http://%s:%d\n\n", port, ip, port)

		addr := fmt.Sprintf(":%d", port)
		if err := app.Listen(addr); err != nil {
			log.Fatal("Failed to start Fiber web project: ", err)
		}
	}()

	select {}
}

func init() {
	runCmd.Flags().BoolP("daemon", "d", false, "Run in daemon mode (persistent execution)")
	runCmd.Flags().IntP("port", "p", 8080, "Port number to run the Gin web project")
	rootCmd.AddCommand(runCmd)
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
