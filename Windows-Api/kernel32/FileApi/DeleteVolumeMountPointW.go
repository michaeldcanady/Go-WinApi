package fileapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

func DeleteVolumeMountPointW(volumeMountPoint string) {
	vmpp, err := syscall.UTF16PtrFromString(volumeMountPoint)
	if err != nil {
		fmt.Println(err)
	}

	ret, _, err := procDeleteVolumeMountPointW.Call(uintptr(unsafe.Pointer(vmpp)))
	if ret == 0 {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
