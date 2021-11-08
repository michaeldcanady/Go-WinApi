package fileapi

import (
	"syscall"
	"unsafe"
)

const (
	//dwAdditionalFlags
	FIND_FIRST_EX_CASE_SENSITIVE       = 1
	FIND_FIRST_EX_LARGE_FETCH          = 2
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY = 4

	//FINDEX_INFO_LEVELS
	FindExInfoStandard     = 0
	FindExInfoBasic        = 1
	FindExInfoMaxInfoLevel = 2

	//FINDEX_SEARCH_OPS
	FindExSearchNameMatch          = 0
	FindExSearchLimitToDirectories = 1
	FindExSearchLimitToDevices     = 2
	FindExSearchMaxSearchOp        = 3

	MAXDWORD = 4294967295
)

var (
	dwFileTypeFlags = map[int64]string{
		0x00000001: "VFT_APP",
		0x00000002: "VFT_DLL",
		0x00000003: "VFT_DRV",
		0x00000004: "VFT_FONT",
		0x00000007: "VFT_STATIC_LIB",
		0x00000000: "VFT_UNKNOWN",
		0x00000005: "VFT_VXD",
	}

	dwFileAttributeFlags = map[int64]string{
		0x20:     "FILE_ATTRIBUTE_ARCHIVE",
		0x800:    "FILE_ATTRIBUTE_COMPRESSED",
		0x40:     "FILE_ATTRIBUTE_DEVICE",
		0x10:     "FILE_ATTRIBUTE_DIRECTORY",
		0x4000:   "FILE_ATTRIBUTE_ENCRYPTED",
		0x2:      "FILE_ATTRIBUTE_HIDDEN",
		0x8000:   "FILE_ATTRIBUTE_INTEGRITY_STREAM",
		0x80:     "FILE_ATTRIBUTE_NORMAL",
		0x2000:   "FILE_ATTRIBUTE_NOT_CONTENT_INDEXED",
		0x20000:  "FILE_ATTRIBUTE_NO_SCRUB_DATA",
		0x1000:   "FILE_ATTRIBUTE_OFFLINE",
		0x1:      "FILE_ATTRIBUTE_READONLY",
		0x400000: "FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS",
		0x40000:  "FILE_ATTRIBUTE_RECALL_ON_OPEN",
		0x400:    "FILE_ATTRIBUTE_REPARSE_POINT",
		0x200:    "FILE_ATTRIBUTE_SPARSE_FILE",
		0x4:      "FILE_ATTRIBUTE_SYSTEM",
		0x100:    "FILE_ATTRIBUTE_TEMPORARY",
		0x10000:  "FILE_ATTRIBUTE_VIRTUAL",
	}

	//Reparse
	dwReparseTag = map[int64]string{
		0x80000009: "IO_REPARSE_TAG_CSV",
		0x80000013: "IO_REPARSE_TAG_DEDUP",
		0x8000000A: "IO_REPARSE_TAG_DFS",
		0x80000012: "IO_REPARSE_TAG_DFSR",
		0xC0000004: "IO_REPARSE_TAG_HSM",
		0x80000006: "IO_REPARSE_TAG_HSM2",
		0xA0000003: "IO_REPARSE_TAG_MOUNT_POINT",
		0x80000014: "IO_REPARSE_TAG_NFS",
		0x80000007: "IO_REPARSE_TAG_SIS",
		0xA000000C: "IO_REPARSE_TAG_SYMLINK",
		0x80000008: "IO_REPARSE_TAG_WIM",
	}
)

func FindFirstFileExW(FileName string, fInfoLevelId int32, fSearchOp int32, dwAdditionalFlags DWORD) (syscall.Handle, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	lpFileName, err := syscall.UTF16PtrFromString(FileName)

	if err != nil {
		return 0, WIN32_FIND_DATAA{}, err
	}

	ret, _, err := findFirstFileExWProc.Call(
		lpFileName,
		uintptr(fInfoLevelId),                    // [in] FINDEX_INFO_LEVELS
		uintptr(unsafe.Pointer(&lpFindFileData)), // [out] LPVOID
		uintptr(fSearchOp),                       // [in] FINDEX_SEARCH_OPS
		uintptr(unsafe.Pointer(nil)),             // [in] LPVOID
		uintptr(dwAdditionalFlags),
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

	if syscall.InvalidHandle == syscall.Handle(ret) {
		return syscall.Handle(0), data, err
	}

	return syscall.Handle(ret), data, nil
}
