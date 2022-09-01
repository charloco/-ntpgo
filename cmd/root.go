package main

import (
	"ntpgo/cmd/install"
	"ntpgo/cmd/start"
	"ntpgo/cmd/sync"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ntpgo",
	Short: "ntp implemented with golang",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Init() {
	rootCmd.AddCommand(install.InstallCmd)
	install.Flags()
	rootCmd.AddCommand(sync.SyncCmd)
	rootCmd.AddCommand(start.StartCmd)
}
func main() {
	Init()
	rootCmd.Execute()
}
