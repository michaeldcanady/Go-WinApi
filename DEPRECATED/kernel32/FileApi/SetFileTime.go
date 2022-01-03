package fileapi

import (
	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func SetFileTime(hFile HANDLE, lpCreationTime, lpLastAccessTime, lpLastWriteTime timezoneapi.FILETIME) error {
	ret, _, err := procSetFileTime.Call(
		hFile.ToUintPtr(),
		lpCreationTime.ToUintPtr(),
		lpLastAccessTime.ToUintPtr(),
		lpLastWriteTime.ToUintPtr(),
	)
	if ret == 0 {
		return err
	}

	return nil
}
