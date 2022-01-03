package timezoneapi

import (
	"syscall"
)

var (
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	fileTimeToSystemTimeProc = kernel32.NewProc("FileTimeToSystemTime")
	SystemTimeToFileTimeProc = kernel32.NewProc("SystemTimeToFileTime")
)
