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
	rootCmd.AddCommand(TcpSendCmd)
	//发送次数
	TcpSendCmd.Flags().IntP("count", "c", 10, "--count")
	TcpSendCmd.Flags().StringP("message", "m", "Hello Net-tools!", "--message")
}

var TcpSendCmd = &cobra.Command{
	Use:   "tcpsend",
	Short: "Send tcp messages to the specified host continuously",
	Long:  "Usage: ./net-tools tcpsend 127.0.0.1 19006 -t 100 -m HelloNet-tools!",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools-2.1 tcpsend 127.0.0.1 19006 -c 100 -m conent ")
		}
		times, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.Fatalln(err)
		}
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatalln(err)
		}
		host := fmt.Sprintf("%s:%s", args[0], args[1])
		if ok := utils.IsIp(args[0]); !ok {
			log.Printf("Please enter the correct IP address:%v", args[0])
			os.Exit(1)
		}
		if err := core.TcpSendMsg(host, message, times); err != nil {
			log.Fatalln(err)
		}
	},
}
