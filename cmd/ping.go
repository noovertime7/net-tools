package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net-tools/core"
)

func init() {
	rootCmd.AddCommand(pingCmd)
	pingCmd.Flags().IntP("count", "c", 10, "--count")
	pingCmd.Flags().IntP("size", "l", 56, "--size")
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Simulate the ping command to send icmp packets to the target host",
	Long:  "net-tools ping 192.168.1.1 -c 10 -l 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools ping 192.168.1.1")
		}
		counts, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.Fatalln(err)
		}
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			log.Fatalln(err)
		}
		core.Ping(args[0], size, counts)
	},
}
