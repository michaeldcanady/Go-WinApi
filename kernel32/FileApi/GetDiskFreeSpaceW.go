package fileapi

import (
	"unsafe"
)

func GetDiskFreeSpaceW(lpDirectoryName string) (lpSectorsPerCluster uint32, lpBytesPerSector uint32, lpNumberOfFreeClusters uint32, lpTotalNumberOfClusters uint32, err error) {

	ret, _, err := procGetDiskFreeSpaceW.Call(
		UintptrFromString(lpDirectoryName),
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
