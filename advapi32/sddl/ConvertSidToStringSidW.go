package sddl

import (
	"fmt"
	"syscall"
	"unsafe"

	winapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api"
)

var (
	advapi32                   = syscall.NewLazyDLL("advapi32.dll")
	procConvertSidToStringSidW = advapi32.NewProc("ConvertSidToStringSidW")
)

func ConvertSidToStringSidW(SID *syscall.SID) (string, error) {

	var StringSid *uint16

	ret, _, err := procConvertSidToStringSidW.Call(
		uintptr(unsafe.Pointer(&SID)),
		uintptr(unsafe.Pointer(&StringSid)),
	)

	fmt.Println(err)

	if ret == 0 {
		return "", err
	}

	defer syscall.LocalFree((syscall.Handle)(unsafe.Pointer(StringSid)))

	return winapi.Utf16PtrToString(StringSid), nil
}
