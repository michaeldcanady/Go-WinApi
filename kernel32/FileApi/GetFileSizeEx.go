package fileapi

import (
	"syscall"
)

func GetFileSizeEx(hFile syscall.Handle) (int64, error) {

	ret, _, err := getFileSizeProc.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0 {
		return 0, err
	}

	return int64(ret), nil
}
