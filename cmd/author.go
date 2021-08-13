package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authorCmd represents the author command
var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "显示项目作者",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Author by 七侠镇法师.")
	},
}

func init() {
	rootCmd.AddCommand(authorCmd)

}
