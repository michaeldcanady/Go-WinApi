package fileapi

import (
	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func FileTimeToLocalFileTime(lpFileTime timezoneapi.FILETIME) (lpLocalFileTime timezoneapi.FILETIME, err error) {

	ret, _, err := procFileTimeToLocalFileTime.Call(
		lpFileTime.ToUintPtr(),
		lpLocalFileTime.ToUintPtr(),
	)
	if ret == 0 {
		return lpLocalFileTime, err
	}

	return lpLocalFileTime, nil
}
