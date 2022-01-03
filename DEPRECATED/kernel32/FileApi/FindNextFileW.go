package fileapi

import (
	"unsafe"
)

//FindNextFileW Continues a file search from a previous call to the FindFirstFile, FindFirstFileEx, or FindFirstFileTransacted functions.
func FindNextFileW(hFindFile HANDLE) (Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextFileW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 {
		return Win32FindDataW{}, err
	}

	return newWin32FindData(lpFindFileData), nil
}
