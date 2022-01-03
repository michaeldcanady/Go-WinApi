package fileapi

import (
	"unsafe"
)

func GetFileAttributesExW(fileName string) (Win32FileAttributeData, error) {

	var lpFileInformation Win32FileAttributeDataA

	ret, _, err := procGetFileAttributesExW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&lpFileInformation)),
	)

	if ret == 0 {
		return newWin32FileAttributeData(lpFileInformation), err
	}

	return newWin32FileAttributeData(lpFileInformation), nil
}
