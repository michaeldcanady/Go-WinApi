package fileapi

import (
	"unsafe"
)

func FindNextStreamW(hFindFile HANDLE) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextStreamW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	data := newWin32FindData(lpFindFileData)

	if ret == 0 {
		return HANDLE(0), data, err
	}

	return HANDLE(ret), data, nil
}
