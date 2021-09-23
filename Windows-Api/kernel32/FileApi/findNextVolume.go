package fileapi

import (
	"syscall"
	"unsafe"
)

func FindNextVolume(handle syscall.Handle) (string, bool, error) {
	const noMoreFiles = 18

	guid := make([]uint16, guidBufLen)

	rc, _, err := findNextVolumeWProc.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if rc == 1 {
		return uint16ToString(guid), true, nil
	}

	if err.(syscall.Errno) == noMoreFiles {
		return "", false, nil
	}
	return "", false, err
}
