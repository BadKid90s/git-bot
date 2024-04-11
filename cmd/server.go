package cmd

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
)

var cfgFile string

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "install an application program",
		Long:  `install an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			err := srv.Install()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall an application program",
		Long:  `uninstall an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			err := srv.Uninstall()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start an application program",
		Long:  `start an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("服务正在启动")
			err := srv.Start()
			if err != nil {
				log.Println("服务启动发生错误")
				log.Fatalln(err)
			}
		},
	}
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop an application program",
		Long:  `stop an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			err := srv.Stop()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	restartCmd = &cobra.Command{
		Use:   "restart",
		Short: "restart an application program",
		Long:  `restart an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			err := srv.Restart()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "status an application program",
		Long:  `status an application program. This is Gitlab-Bot`,
		Run: func(cmd *cobra.Command, args []string) {
			status, err := srv.Status()
			if err != nil {
				log.Println("无法获取服务的状态")
			}
			if status == service.StatusRunning {
				log.Println("服务正在运行中")
			} else if status == service.StatusStopped {
				log.Println("服务已停止")
			} else {
				log.Println("服务状态未知")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(statusCmd)
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	configName := fmt.Sprintf("%s/gitlab-bot.yml", dir)
	log.Println(configName)

	startCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", configName, "config file (default is $HOME/gitlab-bot.yml)")

}
