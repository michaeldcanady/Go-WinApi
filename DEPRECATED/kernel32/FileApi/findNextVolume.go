package fileapi

import (
	"syscall"
	"unsafe"
)

func FindNextVolume(handle HANDLE) (string, bool, error) {
	const noMoreFiles = 18

	guid := make([]uint16, guidBufLen)

	rc, _, err := findNextVolumeWProc.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if rc == 1 {
		return syscall.UTF16ToString(guid), true, nil
	}

	if err.(syscall.Errno) == noMoreFiles {
		return "", false, nil
	}
	return "", false, err
}
