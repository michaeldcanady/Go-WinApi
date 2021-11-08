package fileapi

import (
	"unsafe"
)

func FindFirstFileW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstFileW.Call(
		UintptrFromString(fileName),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	data := newWin32FindData(lpFindFileData)

	if ret == 0 || HANDLE(ret) == INVALID_HANDLE_VALUE {
		return HANDLE(0), data, err
	}

	return HANDLE(ret), data, nil
}
