package winapi

import (
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
	ret, _, err := procDeleteProfile.Call(UserSid, 0, 0)
	switch ret {
	case 0:
		return err
	default:
		return nil
	}
}
