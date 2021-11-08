package fileapi

import (
	"fmt"
	"syscall"
)

func GetFileType(hFile syscall.Handle) (string, error) {

	ret, _, err := getFileTypeProc.Call(uintptr(hFile))

	switch ret {
	case 0x0000:
		if fmt.Sprintf("%s", err) != "The handle is invalid." {
			return "FILE_TYPE_UNKNOWN", nil
		}
	case 0x0001:
		return "FILE_TYPE_DISK", nil
	case 0x0002:
		return "FILE_TYPE_CHAR", nil
	case 0x0003:
		return "FILE_TYPE_PIPE", nil
	case 0x8000:
		return "FILE_TYPE_REMOTE", nil
	}

	return "", err
}
