package fileapi

import (
	"unsafe"
)

//FindFirstStreamW Enumerates the first stream with a ::$DATA stream type in the specified file or directory.
//
//To perform this operation as a transacted operation, use the FindFirstStreamTransactedW function.
func FindFirstStreamW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstStreamW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&lpFindFileData)),
		0,
	)

	if ret == 0 {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}
