package fileapi

import (
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func FileTimeToLocalFileTime(lpFileTime timezoneapi.FILETIME) (lpLocalFileTime timezoneapi.FILETIME, err error) {

	ret, _, err := procFileTimeToLocalFileTime.Call(
		uintptr(unsafe.Pointer(&lpFileTime)),
		uintptr(unsafe.Pointer(&lpLocalFileTime)),
	)
	if ret == 0 {
		return lpLocalFileTime, err
	}

	return lpLocalFileTime, nil
}
