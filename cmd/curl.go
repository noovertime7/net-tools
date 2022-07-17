package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net-tools/core"
)

func init() {
	curlCmd.Flags().StringP("method", "X", "GET", "./net-tools -X POST www.baidu.com")
	curlCmd.Flags().StringP("data", "d", "Hello Net-tools!", "./net-tools -X POST www.baidu.com -d {'msg':'Hello Net-tools!'}")
	rootCmd.AddCommand(curlCmd)
}

var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "net-tools curl",
	Long:  "Usage: ./net-tools curl -X POST http://127.0.0.1:8080 -d {}",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: . /net-tools curl https://www.baidu.com")
		}
		method, err := cmd.Flags().GetString("method")
		if err != nil {
			log.Fatalln(err)
		}
		conent, err := cmd.Flags().GetString("data")
		if err != nil {
			log.Fatalln(err)
		}
		if err := core.CurlCheck(args[0], method, conent); err != nil {
			log.Fatalln(err)
		}
	},
}
