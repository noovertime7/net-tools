package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net-tools/core"
	"net-tools/utils"
	"os"
)

func init() {
	rootCmd.AddCommand(UdptelnetCmd)
	UdptelnetCmd.Flags().IntP("timeout", "t", 5, "-- timeout")
}

var UdptelnetCmd = &cobra.Command{
	Use:   "udptelnet",
	Short: "net-tools udp telnet",
	Long:  "Usage: ./net-tools udptelnet 127.0.0.1 3306",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatalln("You have entered the wrong parameter, Usage: . /net-tools udptelnet 127.0.0.1 3306")
		}
		timeout, err := cmd.Flags().GetInt("timeout")
		if err != nil {
			log.Fatalln(err)
		}
		host := args[0]
		port := args[1]
		if ok := utils.IsIp(host); !ok {
			log.Printf("Please enter the correct IP address:%v", host)
			os.Exit(1)
		}
		if err := core.UdpTelnetChekck(fmt.Sprintf("%v:%v", host, port), timeout); err != nil {
			log.Fatalln(err)
		}
		log.Printf("udptelnet %s connection successful", host)
	},
}
