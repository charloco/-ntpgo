package sync

import (
	"ntpgo/config"
	"os/exec"
	"time"

	"github.com/beevik/ntp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func updateTime(ntpTime string) {
	err := exec.Command("date", "-s", ntpTime).Run()
	if err != nil {
		logrus.Error(err)
	}
}
func sync() {
	for i := 1; i < 3; i++ {
		var ntpTime time.Time
		ntpServer, err := config.GetNtpServer()
		if err == nil {
			ntpTime, err = ntp.Time(ntpServer)
		}
		//如果有错误, 短实际内尝试三次
		if err != nil {
			logrus.Error(err)
			time.Sleep(5 * time.Second)
			continue
		}
		t_str := ntpTime.Local().Format(time.RFC3339)
		logrus.Info(t_str)
		now := time.Now().UTC()
		//如果差距大于2秒, 需要更新
		if now.Sub(ntpTime) > 2*time.Second {
			updateTime(t_str)
		}
		if ntpTime.Sub(now) > 2*time.Second {
			updateTime(t_str)
		}
		break
	}
}

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync time from ntp server and update local time",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			sync()
			time.Sleep(5 * time.Minute)
		}
	},
}
