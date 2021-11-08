package winapi

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	UserenvDLL = windows.NewLazyDLL("Userenv.dll")

	procDeleteProfile = UserenvDLL.NewProc("DeleteProfileW")
)

// Removes user and User directory!!!
func DeleteProfile(SID string) error {
	UserSid := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(SID)))
	ret, _, _ := procDeleteProfile.Call(UserSid, 0, 0)
	switch ret {
	default:
		return fmt.Errorf("Error has appeared, return value %x.\n Please create an issue.", ret)
	}
}
