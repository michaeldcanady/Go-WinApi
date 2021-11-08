package fileapi

import (
	"errors"
	"fmt"
	"syscall"
)

func CreateDirectoryW(pathName string, lpSecurityAttributes interface{}) error {

	if lpSecurityAttributes.(int) != 0 && fmt.Sprintf("%T", lpSecurityAttributes) != "*_SECURITY_ATTRIBUTES" {
		return errors.New("Incompatible lpSecurityAttributes used please use 0 or specify security attributes")
	}

	lpPathName, err := syscall.UTF16PtrFromString(pathName)

	if err != nil {
		return err
	}

	r, _, err := procCreateDirectoryW.Call(
		lpPathName,                          // [in] LPCTSTR
		uintptr(lpSecurityAttributes.(int)), // [in] LPSECURITY_ATT...
	)

	if r != 1 {
		return err
	}

	return nil
}
