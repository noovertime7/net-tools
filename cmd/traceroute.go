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
	Short: "Simulate traceroute command to detect routes",
	Long:  "net-tools traceroute 192.168.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools traceroute 192.168.1.1")
		}
		host := args[0]
		core.Trace(host)
	},
}
