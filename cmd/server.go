package cmd

import (
	"github.com/spf13/cobra"
	"gitlab-bot/internal/core"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start an application program",
	Long:  `start an application program. This is Gitlab-Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		bot := core.NewGitLabBot()
		bot.Run()
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop an application program",
	Long:  `stop an application program. This is Gitlab-Bot`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	startCmd.PersistentFlags().StringVar(&core.CfgFile, "config", "", "config file (default is ./config/application.yaml)")
}
