package fileapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

func GetDriveTypeW(PathName string) (string, error) {

	vmpp, err := syscall.UTF16PtrFromString(PathName)
	if err != nil {
		fmt.Println(err)
	}

	ret, _, err := procGetDriveTypeW.Call(uintptr(unsafe.Pointer(vmpp)))
	if ret == 0 {
		fmt.Println(err)
	}

	switch ret {
	case DRIVE_UNKNOWN:
		return "unknown type", nil
	case DRIVE_NO_ROOT_DIR:
		return "no root dir", nil
	case DRIVE_REMOVABLE:
		return "removable", nil
	case DRIVE_FIXED:
		return "fixed", nil
	case DRIVE_REMOTE:
		return "remote", nil
	case DRIVE_CDROM:
		return "cdrom", nil
	case DRIVE_RAMDISK:
		return "ramdisk", nil
	}
	return "", nil
}
