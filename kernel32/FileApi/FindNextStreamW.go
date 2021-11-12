package fileapi

import (
	"unsafe"
)

//FindNextStreamW Continues a stream search started by a previous call to the FindFirstStreamW function.
func FindNextStreamW(hFindFile HANDLE) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextStreamW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}
