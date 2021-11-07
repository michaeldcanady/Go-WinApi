package fileapi

import (
	"syscall"
	"unsafe"
)

var getLongPathNameWProc = kernel32.NewProc("GetLongPathNameW")

func GetLongPathNameW(lpszShortPath string) (string, error) {
	var bufSize uint32 = syscall.MAX_PATH // 260
	lpszLongPath := make([]uint16, bufSize)

	ret, _, err := getLongPathNameWProc.Call(
		UintptrFromString(&lpszShortPath),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(bufSize),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}
