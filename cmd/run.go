package cmd

import (
	"HerosEvil/app"
	"github.com/spf13/cobra"
	"log"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行此项目",
	Long: `Example:
  HerosEvil.exe run -u test -p 123456 -s h6    //登陆混服6区`,
	RunE: func(cmd *cobra.Command, args []string) error {

		app.LoginStruct.Username, _ = cmd.Flags().GetString("username")
		app.LoginStruct.Password, _ = cmd.Flags().GetString("password")
		tmp, _ := cmd.Flags().GetString("servercode")
		err :=app.ChooseServer(tmp)
		if err != nil {
			return err
		}
		log.Println(app.LoginStruct.Username,app.LoginStruct.Password,app.LoginStruct.QuickLoginUrl,app.LoginStruct.KillNpcUrl)
		app.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&username, "username", "u", "", "-u参数后接登陆账号.")
	runCmd.Flags().StringVarP(&password, "password", "p", "", "-s参数后接登陆密码.")
	runCmd.Flags().StringVarP(&server, "servercode", "s", "", "-s参数后接登陆区服代码.")
	_ = runCmd.MarkFlagRequired("username")
	_ = runCmd.MarkFlagRequired("password")
	_ = runCmd.MarkFlagRequired("servercode")
}

