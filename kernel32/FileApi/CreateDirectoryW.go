package fileapi

import (
	"errors"
	"fmt"
)

var (
	procCreateDirectoryW = kernel32.NewProc("CreateDirectoryW")
)

func CreateDirectoryW(FolderName string, lpSecurityAttributes interface{}) error {

	if lpSecurityAttributes.(int) != 0 && fmt.Sprintf("%T", lpSecurityAttributes) != "*_SECURITY_ATTRIBUTES" {
		return errors.New("Incompatible lpSecurityAttributes used please use 0 or specify security attributes")
	}

	r, _, err := procCreateDirectoryW.Call(
		UintptrFromString(&FolderName),      // [in] LPCTSTR
		uintptr(lpSecurityAttributes.(int)), // [in] LPSECURITY_ATT...
	)

	if r != 1 {
		return err
	}

	return nil
}
