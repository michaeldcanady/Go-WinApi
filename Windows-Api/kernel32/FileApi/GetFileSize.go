package fileapi

import (
	"syscall"
)

var getFileSizeProc = kernel32.NewProc("GetFileSize")

func GetFileSize(hFile syscall.Handle) (int64, error) {

	ret, _, err := getFileSizeProc.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0xFFFFFFFF {
		return 0, err
	}

	return int64(ret), nil
}
