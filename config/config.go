package config

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/goccy/go-yaml"
)

type ntpConfig struct {
	NtpServer string `json:"ntp_server"`
}

func GetNtpServer() (string, error) {
	f, err := os.Open("/etc/ntpgo/ntp.yaml")
	if err != nil {
		return "", err
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	var ntpCfg ntpConfig
	err = yaml.Unmarshal(bs, &ntpCfg)
	if err != nil {
		return "", err
	}
	return ntpCfg.NtpServer, nil
}
func BuildNewConfig(ntpServer string) error {
	ntpCfg := ntpConfig{NtpServer: ntpServer}
	bs, err := yaml.Marshal(ntpCfg)
	if err != nil {
		return err
	}
	err = exec.Command("mkdir", "-p", "/etc/ntpgo").Run()
	if err != nil {
		return err
	}
	f, err := os.Create("/etc/ntpgo/ntp.yaml")
	if err != nil {
		return err
	}
	_, err = f.Write(bs)
	return err
}
