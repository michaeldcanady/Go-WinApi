package fileapi

import (
	"syscall"
	"unsafe"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

var setFileTimeProc = kernel32.NewProc("SetFileTime")

func SetFileTime(hFile syscall.Handle, lpCreationTime, lpLastAccessTime, lpLastWriteTime timezoneapi.FILETIME) error {
	ret, _, err := setFileTimeProc.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpCreationTime)),
		uintptr(unsafe.Pointer(&lpLastAccessTime)),
		uintptr(unsafe.Pointer(&lpLastWriteTime)),
	)
	if ret == 0 {
		return err
	}

	return nil
}
