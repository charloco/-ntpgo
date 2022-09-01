package install

import (
	"embed"
	"ntpgo/config"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mrunalp/fileutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//go:embed ntpgo.service
var FS embed.FS

var ntpServer string

func installService() error {
	bs, err := FS.ReadFile("ntpgo.service")
	if err != nil {
		logrus.Error(err)
		return err
	}
	f, err := os.Create("/etc/systemd/system/ntpgo.service")
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer f.Close()
	_, err = f.Write(bs)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = exec.Command("systemctl", "is-enabled", "ntpgo").Run()
	if err != nil {
		err = exec.Command("systemctl", "enable", "ntpgo").Run()
		if err != nil {
			logrus.WithField("err", err).Error("systemctl enable ntpgo failed")
		}
	}
	return err
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install ntpgo systemd service",
	Run: func(cmd *cobra.Command, args []string) {
		if ntpServer != "" {
			err := config.BuildNewConfig(ntpServer)
			if err != nil {
				logrus.Fatal(err)
			}
		}
		file, err := exec.LookPath(os.Args[0])
		if err != nil {
			logrus.Fatal(err)
		}
		fullpath, err := filepath.Abs(file)
		if err != nil {
			logrus.Fatal(err)
		}
		if fullpath != "/usr/bin/ntpgo" {
			err = fileutils.CopyFile(fullpath, "/usr/bin/ntpgo")
			if err != nil {
				logrus.Fatal(err)
			}
		}
		err = installService()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func Flags() {
	InstallCmd.Flags().StringVar(&ntpServer, "server", "", "ntp server address or domain")
}
