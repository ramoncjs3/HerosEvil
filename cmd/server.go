package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "获取登陆区服代码",
	Long:  `必须运行server命令获取区服代码`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`	"选择区服"     "代码" 
	"官方1区"  ->  "g1"
	"官方2区"  ->  "g2"
	"官方3区"  ->  "g3"
	"混服1区"  ->  "h1"
	"混服2区"  ->  "h2"
	"混服3区"  ->  "h3"
	"混服4区"  ->  "h4"
	"混服5区"  ->  "h5"
	"混服6区"  ->  "h6"
	"混服7区"  ->  "h7"
	"混服8区"  ->  "h8"
	"混服9区"  ->  "h9"
	"混服10区" ->  "h10"`)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
