package fileapi

import (
	"syscall"
	"unsafe"
)

func FindFirstFileW(fileName string) (syscall.Handle, Win32FindDataW, error) {
	
	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstFileW.Call(
		UintptrFromString(fileName),
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

	if syscall.Handle(ret) == syscall.InvalidHandle {
		return syscall.Handle(0), data, err
	}

	return syscall.Handle(ret), data, nil
}
