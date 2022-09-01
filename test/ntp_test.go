package test

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestNTPRequest(t *testing.T) {
	ntpTime, err := ntp.Time("cn.pool.ntp.org")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC()
	if now.Sub(ntpTime) > 10*time.Second {
		t.Fatal("need to update time")
	}
	if ntpTime.Sub(now) > 10*time.Second {
		t.Fatal("need to update time")
	}
}
