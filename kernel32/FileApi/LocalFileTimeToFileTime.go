package fileapi

import (
	"unsafe"
	
	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func LocalFileTimeToFileTime(in timezoneapi.FILETIME) (out timezoneapi.FILETIME, err error) {

	ret, _, err := procLocalFileTimeToFileTime.Call(
		uintptr(unsafe.Pointer(&in)),
		uintptr(unsafe.Pointer(&out)),
	)
	if ret == 0 {
		return out, err
	}
	return out, nil
}
