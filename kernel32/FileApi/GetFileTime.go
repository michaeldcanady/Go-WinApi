package fileapi

import (
	"time"
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func GetFileTime(hFile HANDLE) (time.Time, time.Time, time.Time, error) {

	var dwCreationTime, dwLastAccessTime, dwLastWriteTime timezoneapi.FILETIME

	ret, _, err := procGetFileTime.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&dwCreationTime)),
		uintptr(unsafe.Pointer(&dwLastAccessTime)),
		uintptr(unsafe.Pointer(&dwLastWriteTime)),
	)

	if ret == 0 {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	CreationTime, err := timezoneapi.FileTimeToSystemTime(dwCreationTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}
	LastAccessTime, err := timezoneapi.FileTimeToSystemTime(dwLastAccessTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}
	LastWriteTime, err := timezoneapi.FileTimeToSystemTime(dwLastWriteTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	return CreationTime, LastAccessTime, LastWriteTime, nil
}
