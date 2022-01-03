package fileapi

import (
	"syscall"
	"unsafe"
)

func GetLogicalDriveStringsW() (string, error) {

	buf := make([]uint16, MAX_PATH)

	ret, _, err := procGetLogicalDriveStringsW.Call(
		uintptr(MAX_PATH),
		uintptr(unsafe.Pointer(&buf[0])),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil

}
