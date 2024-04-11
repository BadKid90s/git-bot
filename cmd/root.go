package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitlab-bot",
	Short: "gitlab bot is an extension of gitlab that provides brief automation features",
	Run: func(cmd *cobra.Command, args []string) {
		err := startAction()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
