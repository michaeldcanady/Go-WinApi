package fileapi

import (
	"syscall"
	"unsafe"
)

var writeFileProc = kernel32.NewProc("WriteFile")

func WriteFile(hFile syscall.Handle, data string) error {

	lpBuffer, err := syscall.UTF16FromString(data)
	if err != nil {
		return err
	}

	var buffer uint32

	ret, _, err := writeFileProc.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpBuffer)),
		uintptr(len(lpBuffer)),
		uintptr(unsafe.Pointer(&buffer)),
		0,
	)

	if ret == 0 {
		return err
	}

	return nil
}
