package fileapi

import (
	"syscall"
	"unsafe"
)

func GetLongPathNameW(shortPath string) (string, error) {

	lpszLongPath := make([]uint16, MAX_PATH)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(shortPath),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(MAX_PATH),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}
