package start

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "restart",
	Short: "systemctl restart ntpgo",
	Run: func(cmd *cobra.Command, args []string) {
		err := exec.Command("systemctl", "restart", "ntpgo").Run()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}
