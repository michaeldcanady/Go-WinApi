package fileapi

import (
	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func LocalFileTimeToFileTime(in timezoneapi.FILETIME) (out timezoneapi.FILETIME, err error) {

	ret, _, err := procLocalFileTimeToFileTime.Call(
		in.ToUintPtr(),
		out.ToUintPtr(),
	)
	if ret == 0 {
		return out, err
	}
	return out, nil
}
