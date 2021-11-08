package fileapi

import (
	"unsafe"
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

//FindFirstFileExW Searches a directory for a file or subdirectory with a name and attributes that match those specified.
//For the most basic version of this function, see FindFirstFile.
//To perform this operation as a transacted operation, use the FindFirstFileTransacted function.
func FindFirstFileExW(FileName string, fInfoLevelId int32, fSearchOp int32, dwAdditionalFlags DWORD) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstFileExW.Call(
		UintptrFromString(FileName),
		uintptr(fInfoLevelId),                    // [in] FINDEX_INFO_LEVELS
		uintptr(unsafe.Pointer(&lpFindFileData)), // [out] LPVOID
		uintptr(fSearchOp),                       // [in] FINDEX_SEARCH_OPS
		uintptr(unsafe.Pointer(nil)),             // [in] LPVOID
		uintptr(dwAdditionalFlags),
	)

	handle := HANDLE(ret)

	data := newWin32FindData(lpFindFileData)

	if INVALID_HANDLE_VALUE == handle {
		return handle, data, err
	}

	return handle, data, nil
}
