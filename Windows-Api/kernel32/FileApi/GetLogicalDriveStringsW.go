package fileapi

import (
	"syscall"
	"unsafe"
)

var getLogicalDriveStringsWProc = kernel32.NewProc("GetLogicalDriveStringsW")

func GetLogicalDriveStringsW() (string, error) {

	var bufSize uint32 = syscall.MAX_PATH // 260
	buf := make([]uint16, bufSize)

	ret, _, err := getLogicalDriveStringsWProc.Call(
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&buf[0])),
	)

	if ret == 0 {
		return "", err
	}

	return uint16ToString(buf), nil

}
