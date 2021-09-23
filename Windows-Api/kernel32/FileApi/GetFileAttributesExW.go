package fileapi

import (
	"time"
	"unsafe"

	"github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/timezoneapi"
)

var getFileAttributesExWProc = kernel32.NewProc("GetFileAttributesExW")

type Win32FileAttributeData struct {
	FileAttributes []string
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	FileSize       int
}

type Win32FileAttributeDataA struct {
	FileAttributes uint32
	CreationTime   timezoneapi.FILETIME
	LastAccessTime timezoneapi.FILETIME
	LastWriteTime  timezoneapi.FILETIME
	FileSizeHigh   uint32
	FileSizeLow    uint32
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
		FileAttributes: seperateFlags(data.FileAttributes, dwFileAttributeFlags),
		CreationTime:   CreationTime,
		LastAccessTime: LastAccessTime,
		LastWriteTime:  LastWriteTime,
		FileSize:       highAndLowToSize(data.FileSizeHigh, data.FileSizeLow),
	}
}

func GetFileAttributesExW(lpFileName string) (Win32FileAttributeData, error) {

	var lpFileInformation Win32FileAttributeDataA

	ret, _, err := getFileAttributesExWProc.Call(
		UintptrFromString(&lpFileName),
		0,
		uintptr(unsafe.Pointer(&lpFileInformation)),
	)

	if ret == 0 {
		return newWin32FileAttributeData(lpFileInformation), err
	}

	return newWin32FileAttributeData(lpFileInformation), nil
}
