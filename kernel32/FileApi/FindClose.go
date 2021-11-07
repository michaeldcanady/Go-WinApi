package fileapi

import (
	"syscall"
)

var findCloseProc = kernel32.NewProc("FindClose")

func FindClose(hFindFile syscall.Handle) error {
	ret, _, err := findCloseProc.Call(uintptr(hFindFile))

	if ret == 0 {
		return err
	}

	return nil
}
