package fileapi

import(
	"unsafe"
	"syscall"
)

var(
	findNextFileWProc = kernel32.NewProc("FindNextFileW")
)

func FindNextFileW(hFindFile syscall.Handle) (WIN32_FIND_DATAW, error){
	var lpFindFileData WIN32_FIND_DATAA
	ret, _, err := findNextFileWProc.Call(
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
		return data, err
	}

	return data, nil
}