package fileapi

import (
	"syscall"
	"unsafe"
)

func GetFullPathNameW(lpFileName string) (string, error) {

	var bufSize uint32 = syscall.MAX_PATH // 260
	buf := make([]uint16, bufSize)

	ret, _, err := procGetFullPathNameW.Call(
		UintptrFromString(lpFileName),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&buf[0])),
		0,
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil
}
