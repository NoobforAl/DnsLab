package log

import (
	"fmt"
	"os"
	"runtime"
)

func checkOs() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("C:/Users/%v/Documents/dnslab.log", os.Getenv("username")), nil
	case "linux":
		return "/var/log/dnslab.log", nil
	default:
		return "./dnslab.log", fmt.Errorf("not found os name ! log save './dnslab.log'")
	}
}

func logFilePathSelection() (string, error) {
	return checkOs()
}
