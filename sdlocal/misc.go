package sdlocal

import (
	"os"
	"runtime"
)

func Hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hn
}

func OS() string {
	return runtime.GOOS
}

func Arch() string {
	return runtime.GOARCH
}

func NumCPU() int {
	return runtime.NumCPU()
}
