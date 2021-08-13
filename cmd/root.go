package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var username string
var password string
var server string

var rootCmd = &cobra.Command{
	Use:   "HerosEvil",
	Short: "",
	Long: `提示:
  此项目用于帮助霸刀玩家降低善恶值,
  击杀NPC后不会显示倒地状态,再次击杀需间隔8小时后.
  可能存在封号风险,请三思后谨慎使用,一旦封号与作者无关.

Example:
  HerosEvil.exe server                         //获取登陆区服代码
  HerosEvil.exe run -u test -p 123456 -s h6    //登陆混服6区`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(` _    _                    ______     _ _ 
| |  | |                  |  ____|   (_) |
| |__| | ___ _ __ ___  ___| |____   ___| |
|  __  |/ _ \ '__/ _ \/ __|  __\ \ / / | |
| |  | |  __/ | | (_) \__ \ |___\ V /| | |
|_|  |_|\___|_|  \___/|___/______\_/ |_|_|

			    --help获取帮助
`)
		return nil
	},
	CompletionOptions: struct {
		DisableDefaultCmd   bool
		DisableNoDescFlag   bool
		DisableDescriptions bool
	}{DisableDefaultCmd: true, DisableNoDescFlag: true, DisableDescriptions: true},

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}

func init() {


	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.HerosEvil.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
