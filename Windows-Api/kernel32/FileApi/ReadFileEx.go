package fileapi

import (
	"syscall"
	"unsafe"
)

var readFileExProc = kernel32.NewProc("ReadFileEx")

func ReadFileEx(hFile syscall.Handle) (string, error) {

	var bufSize, err = GetFileSize(hFile)

	if err != nil {
		panic(err)
	}

	lpszLongPath := make([]uint16, bufSize)

	//var ran uint32

	ret, _, err := readFileExProc.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&bufSize)),
		0,
		0,
	)

	if ret == 0 {
		return "", err
	}

	return uint16ToString(lpszLongPath), nil
}
