package fileapi

import (
	"unsafe"
)

var getDiskFreeSpaceWProc = kernel32.NewProc("GetDiskFreeSpaceW")

func GetDiskFreeSpaceW(lpDirectoryName string) (uint32, uint32, uint32, uint32, error) {

	var lpSectorsPerCluster, lpBytesPerSector, lpNumberOfFreeClusters, lpTotalNumberOfClusters uint32

	ret, _, err := getDiskFreeSpaceWProc.Call(
		UintptrFromString(&lpDirectoryName),
		uintptr(unsafe.Pointer(&lpSectorsPerCluster)),
		uintptr(unsafe.Pointer(&lpBytesPerSector)),
		uintptr(unsafe.Pointer(&lpNumberOfFreeClusters)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfClusters)),
	)

	if ret == 0 {
		return 0, 0, 0, 0, err
	}

	return lpSectorsPerCluster, lpBytesPerSector, lpNumberOfFreeClusters, lpTotalNumberOfClusters, nil
}
