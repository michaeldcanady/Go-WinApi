package fileapi

import (
	"syscall"
)

func CreateFile2(fileName string, dwDesiredAccess, dwShareMode, dwCreationDisposition DWORD) (syscall.Handle, error) {

	lpFileName, err := syscall.UTF16PtrFromString(fileName)

	if err != nil {
		return 0, nil
	}

	ret, _, err := createFile2Proc.Call(
		lpFileName,
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(dwCreationDisposition),
		0,
	)

	if ret == 18446744073709551615 {
		return syscall.Handle(ret), err
	}

	return syscall.Handle(ret), nil
}
