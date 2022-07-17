package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net-tools/core"
)

func init() {
	rootCmd.AddCommand(traceRouteCmd)
}

var traceRouteCmd = &cobra.Command{
	Use:   "traceroute",
	Short: "net-tools traceroute",
	Long:  "net-tools traceroute",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools traceroute www.baidu.com")
		}
		host := args[0]
		core.Trace(host)
	},
}
