package fileapi

import (
	"syscall"
	"unsafe"
)

func WriteFile(hFile HANDLE, data string) error {

	lpBuffer, err := syscall.UTF16FromString(data)
	if err != nil {
		return err
	}

	var buffer uint32

	ret, _, err := procWriteFile.Call(
		hFile.ToUintPtr(),
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
