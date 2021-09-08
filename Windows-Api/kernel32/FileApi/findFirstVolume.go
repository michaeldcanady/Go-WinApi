package fileapi

import "unsafe"

func findFirstVolume() (uintptr, []uint16, error) {
	const invalidHandleValue = ^uintptr(0)

	guid := make([]uint16, guidBufLen)

	handle, _, err := findFirstVolumeWProc.Call(
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if handle == invalidHandleValue {
		return invalidHandleValue, nil, err
	}

	return handle, guid, nil
}
