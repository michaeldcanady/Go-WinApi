package fileapi

import (
	"syscall"
	"unsafe"
)

var getShortPathNameWProc = kernel32.NewProc("GetShortPathNameW")

func GetShortPathNameW(lpszLongPath string) (string, error) {
	var bufSize uint32 = syscall.MAX_PATH // 260
	lpszShortPath := make([]uint16, bufSize)

	ret, _, err := getLongPathNameWProc.Call(
		UintptrFromString(&lpszLongPath),
		uintptr(unsafe.Pointer(&lpszShortPath[0])),
		uintptr(bufSize),
	)

	if ret == 0 {
		return "", err
	}

	return uint16ToString(lpszShortPath), nil
}
