package fileapi

import (
	"strings"
	"syscall"
	"unsafe"
)

func GetVolumeNameForVolumeMountPointW(volumeMountPoint string) (string, error) {

	if len(volumeMountPoint) == 0 {
		return "", syscall.EINVAL
	}
	if !strings.HasSuffix(volumeMountPoint, "\\") {
		volumeMountPoint = volumeMountPoint + "\\"
	}

	vmpp, err := syscall.UTF16PtrFromString(volumeMountPoint)
	if err != nil {
		return "", err
	}

	var vnBuffer [MaxVolumeNameLength]uint16
	p0 := &vnBuffer[0]

	re, _, err := procGetVolumeNameForVolumeMountPointW.Call(
		uintptr(unsafe.Pointer(vmpp)),
		uintptr(unsafe.Pointer(p0)),
		uintptr(MaxVolumeNameLength),
	)
	if re == 0 {
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(vnBuffer[:]), nil
}
