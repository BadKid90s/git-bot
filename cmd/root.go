package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab-bot/internal/core"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitlab-bot",
	Short: "gitlab bot is an extension of gitlab that provides brief automation features",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)

	rootCmd.PersistentFlags().StringP("author", "a", "WangRuiYu", "123")
	rootCmd.PersistentFlags().StringP("version", "b", "YOUR NAME", "456")

	cobra.OnInitialize(core.InitConfig)
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}
