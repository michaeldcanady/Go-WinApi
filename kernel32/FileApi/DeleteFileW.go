package fileapi

import "syscall"

var (
	procDeleteFileW = kernel32.NewProc("DeleteFileW")
)

func DeleteFileW(FileName string) error {
	ret, _, err := procDeleteFileW.Call(syscall.UTF16FromString(FileName))

	if ret == 0 {
		return err
	}

	return nil
}
