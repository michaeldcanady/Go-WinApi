package fileapi

import (
	"errors"
	"fmt"
)

func CreateDirectoryW(pathName string, lpSecurityAttributes interface{}) error {


	//TODO find away to remove fmt and errors
	if lpSecurityAttributes.(int) != 0 && fmt.Sprintf("%T", lpSecurityAttributes) != "*_SECURITY_ATTRIBUTES" {
		return errors.New("Incompatible lpSecurityAttributes used please use 0 or specify security attributes")
	}

	r, _, err := procCreateDirectoryW.Call(
		UintptrFromString(pathName),                          // [in] LPCTSTR
		uintptr(lpSecurityAttributes.(int)), // [in] LPSECURITY_ATT...
	)

	if r != 1 {
		return err
	}

	return nil
}
