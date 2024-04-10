package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is Gitlab-Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gitlab-Bot v0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
