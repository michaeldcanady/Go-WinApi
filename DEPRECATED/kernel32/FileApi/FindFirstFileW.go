package fileapi

import (
	"unsafe"
)

//FindFirstFileW Searches a directory for a file or subdirectory with a name that matches a specific name (or partial name if wildcards are used).
//
//To specify additional attributes to use in a search, use the FindFirstFileEx function.
//
//To perform this operation as a transacted operation, use the FindFirstFileTransacted function.
func FindFirstFileW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstFileW.Call(
		UintptrFromString(fileName),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 || HANDLE(ret) == INVALID_HANDLE_VALUE {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}
