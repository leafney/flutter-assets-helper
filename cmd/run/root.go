/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     flutter-assets-helper
 * @Date:        2024-02-05 22:57
 * @Description:
 */

package run

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fah",
	Short: "Fah is the flutter resource management assistant",
	Long: `Fah is the flutter resource management assistant, 
	which makes it easier and more convenient to manage 
	the images, colors, fonts, plugins and other resources 
	used in flutter project development.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fah start...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
