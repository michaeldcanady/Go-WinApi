package fileapi

import (
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

func FileTimeToLocalFileTime(in timezoneapi.FILETIME) (timezoneapi.FILETIME, error) {
	var out timezoneapi.FILETIME
	ret, _, err := fileTimeToLocalFileTimeProc.Call(
		uintptr(unsafe.Pointer(&in)),
		uintptr(unsafe.Pointer(&out)),
	)
	if ret == 0 {
		return out, err
	}
	return out, nil
}
