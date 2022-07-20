package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "net-tools version",
	Long:  "get net-tools version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net-tools v2.2 ")
	},
}
