package fileapi

import (
	"fmt"
	"syscall"
	"time"

	timezoneapi "github.com/michaeldcanady/Go-WinApi/kernel32/TimezoneApi"
)

type WIN32_FIND_DATAA struct {
	dwFileAttributes   uint32
	ftCreationTime     timezoneapi.FILETIME
	ftLastAccessTime   timezoneapi.FILETIME
	ftLastWriteTime    timezoneapi.FILETIME
	nFileSizeHigh      uint32
	nFileSizeLow       uint32
	dwReserved0        uint32
	dwReserved1        uint32
	cFileName          []uint16
	cAlternateFileName []uint16
	dwFileType         uint32
	dwCreatorType      uint32
	wFinderFlags       uint32
}

type Win32FindDataW struct {
	dwFileAttributes   []string
	ftCreationTime     time.Time
	ftLastAccessTime   time.Time
	ftLastWriteTime    time.Time
	nFileSizeHigh      int
	dwReserved0        []string
	dwReserved1        uint32
	cFileName          string
	cAlternateFileName string
	dwFileType         []string
	dwCreatorType      uint32
	wFinderFlags       uint32
}

func newWin32FindData(oldObject WIN32_FIND_DATAA) (data Win32FindDataW) {

	CreationTime, err := timezoneapi.FileTimeToSystemTime(oldObject.ftCreationTime)
	if err != nil {
		fmt.Println(err)
	}
	LastAccessTime, err := timezoneapi.FileTimeToSystemTime(oldObject.ftLastAccessTime)
	if err != nil {
		fmt.Println(err)
	}
	LastWriteTime, err := timezoneapi.FileTimeToSystemTime(oldObject.ftLastWriteTime)
	if err != nil {
		fmt.Println(err)
	}

	data = Win32FindDataW{
		dwFileAttributes:   SeperateFlags(oldObject.dwFileAttributes, dwFileAttributeFlags),
		ftCreationTime:     CreationTime,
		ftLastAccessTime:   LastAccessTime,
		ftLastWriteTime:    LastWriteTime,
		nFileSizeHigh:      highAndLowToSize(oldObject.nFileSizeHigh, oldObject.nFileSizeLow),
		dwReserved0:        SeperateFlags(oldObject.dwReserved0, dwReparseTag),
		dwReserved1:        oldObject.dwReserved1,
		cFileName:          syscall.UTF16ToString(oldObject.cFileName),
		cAlternateFileName: syscall.UTF16ToString(oldObject.cAlternateFileName),
		dwFileType:         SeperateFlags(oldObject.dwFileType, volumeFlags),
		dwCreatorType:      oldObject.dwCreatorType,
		wFinderFlags:       oldObject.wFinderFlags,
	}

	return
}

type Win32FileAttributeData struct {
	FileAttributes []string
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	FileSize       int
}

func newWin32FileAttributeData(data Win32FileAttributeDataA) Win32FileAttributeData {
	CreationTime, err := timezoneapi.FileTimeToSystemTime(data.CreationTime)
	if err != nil {
		panic(err)
	}
	LastAccessTime, err := timezoneapi.FileTimeToSystemTime(data.LastAccessTime)
	if err != nil {
		panic(err)
	}
	LastWriteTime, err := timezoneapi.FileTimeToSystemTime(data.LastWriteTime)
	if err != nil {
		panic(err)
	}
	return Win32FileAttributeData{
		FileAttributes: SeperateFlags(data.FileAttributes, dwFileAttributeFlags),
		CreationTime:   CreationTime,
		LastAccessTime: LastAccessTime,
		LastWriteTime:  LastWriteTime,
		FileSize:       highAndLowToSize(data.FileSizeHigh, data.FileSizeLow),
	}
}

type Win32FileAttributeDataA struct {
	FileAttributes uint32
	CreationTime   timezoneapi.FILETIME
	LastAccessTime timezoneapi.FILETIME
	LastWriteTime  timezoneapi.FILETIME
	FileSizeHigh   uint32
	FileSizeLow    uint32
}

type Volume struct {
	PathName                 string
	VolumeLabel              string
	nVolumeNameSize          uint32
	SerialNumber             uint32
	LpMaximumComponentLength uint32
	SystemFlags              []string
	FileSystem               string
	nFileSystemNameSize      uint32
}

func newVolume(lpRootPathName string, lpVolumeNameBuffer []uint16, nVolumeNameSize, lpVolumeSerialNumber, lpMaximumComponentLength uint32, lpFileSystemFlags uint32, lpFileSystemNameBuffer []uint16, nFileSystemNameSize uint32) Volume {
	label := syscall.UTF16ToString(lpVolumeNameBuffer)
	return Volume{
		PathName:                 lpRootPathName,
		VolumeLabel:              label,
		nVolumeNameSize:          nVolumeNameSize,
		SerialNumber:             lpVolumeSerialNumber,
		LpMaximumComponentLength: lpMaximumComponentLength,
		SystemFlags:              SeperateFlags(lpFileSystemFlags, volumeFlags),
		FileSystem:               syscall.UTF16ToString(lpFileSystemNameBuffer),
		nFileSystemNameSize:      nFileSystemNameSize,
	}
}

type _SECURITY_ATTRIBUTES struct {
	nLength              DWORD
	lpSecurityDescriptor LPVOID
	bInheritHandle       bool
}

func NewSecurityAttribtute(nLength uint32, lpSecurityDescriptor LPVOID, bInheritHandle bool) _SECURITY_ATTRIBUTES {

	return _SECURITY_ATTRIBUTES{
		nLength:              DWORD(nLength),
		lpSecurityDescriptor: lpSecurityDescriptor,
		bInheritHandle:       bInheritHandle,
	}
}
