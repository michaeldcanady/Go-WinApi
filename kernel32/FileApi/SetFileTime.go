package fileapi

import (
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

func SetFileTime(hFile HANDLE, lpCreationTime, lpLastAccessTime, lpLastWriteTime timezoneapi.FILETIME) error {
	ret, _, err := setFileTimeProc.Call(
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
