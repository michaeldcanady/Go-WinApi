package fileapi

import (
	"syscall"
	"unsafe"
)

func FindFirstVolume() (syscall.Handle, string, error) {
	const invalidHandleValue = ^uintptr(0)

	guid := make([]uint16, guidBufLen)

	handle, _, err := findFirstVolumeWProc.Call(
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if handle == invalidHandleValue {
		return syscall.InvalidHandle, "", err
	}

	return syscall.Handle(handle), uint16ToString(guid), nil
}
