package fileapi

import (
	"syscall"
)

func FindClose(hFindFile syscall.Handle) error {
	ret, _, err := findCloseProc.Call(uintptr(hFindFile))

	if ret == 0 {
		return err
	}

	return nil
}
