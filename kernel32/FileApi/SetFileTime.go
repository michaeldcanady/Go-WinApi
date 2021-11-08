package fileapi

import (
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

func SetFileTime(hFile HANDLE, lpCreationTime, lpLastAccessTime, lpLastWriteTime timezoneapi.FILETIME) error {
	ret, _, err := procSetFileTime.Call(
		hFile.toUTF16Ptr(),
		uintptr(unsafe.Pointer(&lpCreationTime)),
		uintptr(unsafe.Pointer(&lpLastAccessTime)),
		uintptr(unsafe.Pointer(&lpLastWriteTime)),
	)
	if ret == 0 {
		return err
	}

	return nil
}
