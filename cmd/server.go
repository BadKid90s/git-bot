package cmd

import (
	"github.com/spf13/cobra"
	"gitlab-bot/internal/core"
)

var bot *core.GitLabBot
var cfgFile string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start an application program",
	Long:  `start an application program. This is Gitlab-Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		bot = core.NewGitLabBot()
		bot.SetConfig(&cfgFile)
		bot.Start()
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop an application program",
	Long:  `stop an application program. This is Gitlab-Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		bot.Stop()
	},
}

func init() {

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)

	startCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "./config/application.yml", "config file (default is ./config/application.yaml)")

}
