package fileapi

import (
	"syscall"
)

var createFile2Proc = kernel32.NewProc("CreateFile2")

func CreateFile2(lpFileName string, dwDesiredAccess, dwShareMode, dwCreationDisposition DWORD) (syscall.Handle, error) {
	ret, _, err := createFile2Proc.Call(
		syscall.UTF16FromString(lpFileName),
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
