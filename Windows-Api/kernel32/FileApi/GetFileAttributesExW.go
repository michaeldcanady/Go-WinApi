package fileapi

import (
	"unsafe"
)

var getFileAttributesExWProc = kernel32.NewProc("GetFileAttributesExW")

func GetFileAttributesExW(lpFileName string) (Win32FileAttributeData, error) {

	var lpFileInformation Win32FileAttributeDataA

	ret, _, err := getFileAttributesExWProc.Call(
		UintptrFromString(&lpFileName),
		0,
		uintptr(unsafe.Pointer(&lpFileInformation)),
	)

	if ret == 0 {
		return newWin32FileAttributeData(lpFileInformation), err
	}

	return newWin32FileAttributeData(lpFileInformation), nil
}
