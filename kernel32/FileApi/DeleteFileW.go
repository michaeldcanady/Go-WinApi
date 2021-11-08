package fileapi

import "syscall"

func DeleteFileW(FileName string) error {
	fileNameInt, err := syscall.UTF16PtrFromString(FileName)

	if err != nil {
		return err
	}

	ret, _, err := procDeleteFileW.Call(fileNameInt)

	if ret == 0 {
		return err
	}

	return nil
}
