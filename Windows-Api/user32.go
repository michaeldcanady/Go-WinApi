package winapi

import (
	"syscall"
	"unsafe"
)

const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL

	_ unsafe.Pointer

	// Points to the desired dll or 'WINDOWS API'
	user32DLL, _ = syscall.LoadDLL("user32.dll")

	procSystemParamInfo, _ = user32DLL.FindProc("SystemParametersInfoW")
)

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}

	return e
}
