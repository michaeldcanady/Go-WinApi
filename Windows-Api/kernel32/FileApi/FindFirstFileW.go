package fileapi

import(
	"unsafe"
	"syscall"
)

var findFirstFileWProc = kernel32.NewProc("FindFirstFileW")

func FindFirstFileW(lpFileName string)(syscall.Handle, WIN32_FIND_DATAW, error){
	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := findFirstFileWProc.Call(
		UintptrFromString(&lpFileName),
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
		return syscall.Handle(0),data, err
	}

	return syscall.Handle(ret),data, nil
}