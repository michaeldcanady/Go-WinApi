package handleapi

import "syscall"

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
)
