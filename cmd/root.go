package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "net-tools",
	Short: "A tool to check network connectivity inside containers",
	Long: `A tool to detect network connectivity within the container,
and support detection of middleware such as redis, mysql,traceroute, etc., 
please use net-tools help for detailed usage`,
}

func init() {
	telnetCmd.Flags().IntP("timeout", "t", 5, "-- timeout")
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
