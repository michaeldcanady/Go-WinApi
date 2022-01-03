package fileapi

import (
	"syscall"
	"unsafe"
)

//FindFirstVolume Retrieves the name of a volume on a computer. FindFirstVolume is used to begin scanning the volumes of a computer.
func FindFirstVolume() (HANDLE, string, error) {
	const invalidHandleValue = ^uintptr(0)

	guid := make([]uint16, guidBufLen)

	handle, _, err := findFirstVolumeWProc.Call(
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if handle == invalidHandleValue {
		return INVALID_HANDLE_VALUE, "", err
	}

	return HANDLE(handle), syscall.UTF16ToString(guid), nil
}
