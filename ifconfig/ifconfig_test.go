package ifconfig

import (
	"testing"
	"fmt"
	"regexp"
)

func TestGetPublicIP(t *testing.T) {
	publicIP := GetPublicIP()
	matched, err := regexp.Match(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`, []byte(publicIP))
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error(fmt.Sprintf("%s not a ip", publicIP))
	}

}
