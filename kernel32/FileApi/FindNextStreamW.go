package fileapi

import (
	"syscall"
	"unsafe"
)

func FindNextStreamW(hFindFile HANDLE) (syscall.Handle, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextStreamW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	data := newWin32FindData(
		lpFindFileData.dwFileAttributes,
		lpFindFileData.ftCreationTime,
		lpFindFileData.ftLastAccessTime,
		lpFindFileData.ftLastWriteTime,
		lpFindFileData.nFileSizeHigh,
		lpFindFileData.nFileSizeLow,
		lpFindFileData.dwReserved0,
		lpFindFileData.dwReserved1,
		lpFindFileData.cFileName,
		lpFindFileData.cAlternateFileName,
		lpFindFileData.dwFileType,
		lpFindFileData.dwCreatorType,
		lpFindFileData.wFinderFlags,
	)

	if ret == 0 {
		return syscall.Handle(0), data, err
	}

	return syscall.Handle(ret), data, nil
}
