package fileapi

import(
	"fmt"
	"unsafe"
	"syscall"
)

var(
	findFirstFileExWProc = kernel32.NewProc("FindFirstFileExW")
)

const(
	//dwAdditionalFlags
	FIND_FIRST_EX_CASE_SENSITIVE = 1
	FIND_FIRST_EX_LARGE_FETCH = 2
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY = 4

	//FINDEX_INFO_LEVELS
	FindExInfoStandard = 0
	FindExInfoBasic = 1
	FindExInfoMaxInfoLevel = 2	

	//FINDEX_SEARCH_OPS
	FindExSearchNameMatch = 0
	FindExSearchLimitToDirectories = 1
	FindExSearchLimitToDevices = 2
	FindExSearchMaxSearchOp = 3
)

type WCHAR uint16

type FILETIME struct {
	DwLowDateTime  DWORD
	DwHighDateTime DWORD
}

type WIN32_FIND_DATAW struct {
	dwFileAttributes DWORD
	ftCreationTime FILETIME
	ftLastAccessTime FILETIME
	ftLastWriteTime FILETIME
	nFileSizeHigh DWORD
	nFileSizeLow DWORD
	dwReserved0 DWORD
	dwReserved1 DWORD
	cFileName WCHAR
	cAlternateFileName WCHAR
	dwFileType DWORD
	dwCreatorType DWORD
	wFinderFlags DWORD
  }


func FindFirstFileExW(FileName string, fInfoLevelId  int32, fSearchOp  int32, dwAdditionalFlags DWORD) (syscall.Handle, error){

	var lpFindFileData WIN32_FIND_DATAW

	ret, _, err := findFirstFileExWProc.Call(
		UintptrFromString(&FileName),
		uintptr(fInfoLevelId),                   // [in] FINDEX_INFO_LEVELS
		uintptr(unsafe.Pointer(&lpFindFileData)), // [out] LPVOID
		uintptr(fSearchOp),                      // [in] FINDEX_SEARCH_OPS
		uintptr(unsafe.Pointer(nil)), // [in] LPVOID
		uintptr(dwAdditionalFlags), 
	)

	fmt.Println(lpFindFileData)

	if(syscall.InvalidHandle == syscall.Handle(ret)){
		return syscall.Handle(0), err
	}

	return syscall.Handle(ret), nil
}