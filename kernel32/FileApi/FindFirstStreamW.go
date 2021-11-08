package fileapi

import (
	"unsafe"
)

func FindFirstStreamW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstStreamW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&lpFindFileData)),
		0,
	)

	data := newWin32FindData(lpFindFileData)

	if ret == 0 {
		return HANDLE(0), data, err
	}

	return HANDLE(ret), data, nil
}
