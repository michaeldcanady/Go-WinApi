package fileapi

import (
	"syscall"
	"unsafe"
)

func GetShortPathNameW(longPath string) (string, error) {

	lpszShortPath := make([]uint16, MAX_PATH)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(longPath),
		uintptr(unsafe.Pointer(&lpszShortPath[0])),
		uintptr(MAX_PATH),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszShortPath), nil
}
