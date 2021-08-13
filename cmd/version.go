package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
