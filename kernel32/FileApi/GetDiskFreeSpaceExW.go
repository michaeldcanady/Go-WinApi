package fileapi

import (
	"unsafe"
)

func GetDiskFreeSpaceExW(lpDirectoryName string) (int64, int64, int64, error) {

	var lpFreeBytesAvailableToCaller, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes int64

	ret, _, err := getDiskFreeSpaceExWProc.Call(
		UintptrFromString(&lpDirectoryName),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailableToCaller)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)),
	)
	if ret == 0 {
		return 0, 0, 0, err
	}

	return lpFreeBytesAvailableToCaller, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes, nil
}
