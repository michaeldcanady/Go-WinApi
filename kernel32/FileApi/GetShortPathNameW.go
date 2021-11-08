package fileapi

import (
	"syscall"
	"unsafe"
)

func GetShortPathNameW(longPath string) (string, error) {
	var bufSize uint32 = syscall.MAX_PATH // 260
	lpszShortPath := make([]uint16, bufSize)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(longPath),
		uintptr(unsafe.Pointer(&lpszShortPath[0])),
		uintptr(bufSize),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszShortPath), nil
}
