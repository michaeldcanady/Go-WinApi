package fileapi

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

var (
	findFirstFileExWProc = kernel32.NewProc("FindFirstFileExW")
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
)

type WCHAR uint16

const (
	FILE_ATTRIBUTE_ARCHIVE               = 32
	FILE_ATTRIBUTE_COMPRESSED            = 2048
	FILE_ATTRIBUTE_DEVICE                = 64
	FILE_ATTRIBUTE_DIRECTORY             = 16
	FILE_ATTRIBUTE_ENCRYPTED             = 16384
	FILE_ATTRIBUTE_HIDDEN                = 2
	FILE_ATTRIBUTE_INTEGRITY_STREAM      = 32768
	FILE_ATTRIBUTE_NORMAL                = 128
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   = 8192
	FILE_ATTRIBUTE_NO_SCRUB_DATA         = 131072
	FILE_ATTRIBUTE_OFFLINE               = 4096
	FILE_ATTRIBUTE_READONLY              = 1
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS = 4194304
	FILE_ATTRIBUTE_RECALL_ON_OPEN        = 262144
	FILE_ATTRIBUTE_REPARSE_POINT         = 1024
	FILE_ATTRIBUTE_SPARSE_FILE           = 512
	FILE_ATTRIBUTE_SYSTEM                = 4
	FILE_ATTRIBUTE_TEMPORARY             = 256
	FILE_ATTRIBUTE_VIRTUAL               = 65536
	MAXDWORD                             = 4294967295
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
)

const (
	//Reparse
	IO_REPARSE_TAG_CSV         = 0x80000009
	IO_REPARSE_TAG_DEDUP       = 0x80000013
	IO_REPARSE_TAG_DFS         = 0x8000000A
	IO_REPARSE_TAG_DFSR        = 0x80000012
	IO_REPARSE_TAG_HSM         = 0xC0000004
	IO_REPARSE_TAG_HSM2        = 0x80000006
	IO_REPARSE_TAG_MOUNT_POINT = 0xA0000003
	IO_REPARSE_TAG_NFS         = 0x80000014
	IO_REPARSE_TAG_SIS         = 0x80000007
	IO_REPARSE_TAG_SYMLINK     = 0xA000000C
	IO_REPARSE_TAG_WIM         = 0x80000008
)

type WIN32_FIND_DATAA struct {
	dwFileAttributes   DWORD
	ftCreationTime     timezoneapi.FILETIME
	ftLastAccessTime   timezoneapi.FILETIME
	ftLastWriteTime    timezoneapi.FILETIME
	nFileSizeHigh      DWORD
	nFileSizeLow       DWORD
	dwReserved0        DWORD
	dwReserved1        DWORD
	cFileName          [260]WCHAR
	cAlternateFileName [14]WCHAR
	dwFileType         DWORD
	dwCreatorType      DWORD
	wFinderFlags       DWORD
}

type Win32FindDataW struct {
	dwFileAttributes   []string
	ftCreationTime     time.Time
	ftLastAccessTime   time.Time
	ftLastWriteTime    time.Time
	nFileSizeHigh      int
	dwReserved0        DWORD
	dwReserved1        DWORD
	cFileName          string
	cAlternateFileName string
	dwFileType         []string
	dwCreatorType      DWORD
	wFinderFlags       DWORD
}

func newWin32FindData(dwFileAttributes DWORD, ftCreationTime, ftLastAccessTime, ftLastWriteTime timezoneapi.FILETIME, nFileSizeHigh, nFileSizeLow, dwReserved0, dwReserved1 DWORD, cFileName [260]WCHAR, cAlternateFileName [14]WCHAR, dwFileType DWORD, dwCreatorType DWORD, wFinderFlags DWORD) (data Win32FindDataW) {

	intFileAttributesFlags := seperateFlags(uint32(dwFileAttributes), dwFileTypeFlags)

	intFileTypeFlags := seperateFlags(uint32(dwFileType), volumeFlags)

	CreationTime, err := timezoneapi.FileTimeToSystemTime(ftCreationTime)
	if err != nil {
		fmt.Println(err)
	}
	LastAccessTime, err := timezoneapi.FileTimeToSystemTime(ftLastAccessTime)
	if err != nil {
		fmt.Println(err)
	}
	LastWriteTime, err := timezoneapi.FileTimeToSystemTime(ftLastWriteTime)
	if err != nil {
		fmt.Println(err)
	}

	data = Win32FindDataW{
		dwFileAttributes:   intFileAttributesFlags,
		ftCreationTime:     CreationTime,
		ftLastAccessTime:   LastAccessTime,
		ftLastWriteTime:    LastWriteTime,
		nFileSizeHigh:      (int(nFileSizeHigh) * (MAXDWORD + 1)) + int(nFileSizeLow),
		dwReserved0:        dwReserved0,
		dwReserved1:        dwReserved1,
		cFileName:          uint16ToString(cFileName),
		cAlternateFileName: uint16ToString1(cAlternateFileName),
		dwFileType:         intFileTypeFlags,
		dwCreatorType:      dwCreatorType,
		wFinderFlags:       wFinderFlags,
	}

	return
}

func uint16ToString(input [260]WCHAR) (output string) {
	for _, in := range input {
		if in != 0 {
			output += string(in)
		}
	}
	return
}

func uint16ToString1(input [14]WCHAR) (output string) {
	for _, in := range input {
		output += string(in)
	}
	return
}

func FindFirstFileExW(FileName string, fInfoLevelId int32, fSearchOp int32, dwAdditionalFlags DWORD) (syscall.Handle, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := findFirstFileExWProc.Call(
		UintptrFromString(&FileName),
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
