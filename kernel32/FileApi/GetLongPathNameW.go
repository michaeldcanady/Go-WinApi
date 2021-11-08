package fileapi

import (
	"syscall"
	"unsafe"
)

func GetLongPathNameW(shortPath string) (string, error) {
	var bufSize uint32 = syscall.MAX_PATH // 260
	lpszLongPath := make([]uint16, bufSize)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(shortPath),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(bufSize),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}
