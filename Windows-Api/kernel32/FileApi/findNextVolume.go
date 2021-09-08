package fileapi

import (
	"syscall"
	"unsafe"
)

func findNextVolume(handle uintptr) ([]uint16, bool, error) {
	const noMoreFiles = 18

	guid := make([]uint16, guidBufLen)

	rc, _, err := findNextVolumeWProc.Call(
		handle,
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if rc == 1 {
		return guid, true, nil
	}

	if err.(syscall.Errno) == noMoreFiles {
		return nil, false, nil
	}
	return nil, false, err
}
