package fileapi

import (
	timezoneapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

func SetFileTime(hFile HANDLE, lpCreationTime, lpLastAccessTime, lpLastWriteTime timezoneapi.FILETIME) error {

	ret, _, err := setFileTimeProc.Call(
		hFile.toUTF16Ptr(),
		lpCreationTime.toUTF16Ptr(),
		lpLastAccessTime.toUTF16Ptr(),
		lpLastWriteTime.toUTF16Ptr(),
	)
	if ret == 0 {
		return err
	}

	return nil
}
