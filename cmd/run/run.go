/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-05 22:48
 * @Description:
 */

package run

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the Fah web project",
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		if daemon {
			// Start Fiber web project in background (persistent execution)
			go StartWeb()
		} else {
			// Start Fiber web project and open a web page
			StartWeb()
		}
	},
}

func StartWeb() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal("Failed to start Fiber web project: ", err)
		}
	}()

	select {}
}

func init() {
	runCmd.Flags().BoolP("daemon", "d", false, "Run in daemon mode (persistent execution)")
	rootCmd.AddCommand(runCmd)
}
