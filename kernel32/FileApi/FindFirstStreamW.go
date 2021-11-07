package fileapi

import (
	
	"syscall"
	"unsafe"
)

var findFirstStreamWProc = kernel32.NewProc("FindFirstStreamW")

func FindFirstStreamW(lpFileName string) (syscall.Handle, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := findFirstStreamWProc.Call(
		UintptrFromString(&lpFileName),
		0,
		uintptr(unsafe.Pointer(&lpFindFileData)),
		0,
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
