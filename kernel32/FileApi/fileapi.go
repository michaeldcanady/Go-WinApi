package fileapi

import (
	"fmt"
	"syscall"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procCreateFileW                       = kernel32.NewProc("CreateFileW")
	procCreateDirectoryW                  = kernel32.NewProc("CreateDirectoryW")
	areFileApisANSIProc                   = kernel32.NewProc("AreFileApisANSI")
	findFirstVolumeWProc                  = kernel32.NewProc("FindFirstVolumeW")
	findNextVolumeWProc                   = kernel32.NewProc("FindNextVolumeW")
	findVolumeCloseProc                   = kernel32.NewProc("FindVolumeClose")
	getVolumePathNamesForVolumeNameWProc  = kernel32.NewProc("GetVolumePathNamesForVolumeNameW")
	procGetVolumeNameForVolumeMountPointW = kernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	getDriveTypeWProc                     = kernel32.NewProc("GetDriveTypeW")
	procGetVolumeInformationW             = kernel32.NewProc("GetVolumeInformationW")
	procGetLogicalDrives                  = kernel32.NewProc("GetLogicalDrives")
	procDeleteVolumeMountPointW           = kernel32.NewProc("DeleteVolumeMountPointW")
	procGetDriveTypeW                     = kernel32.NewProc("GetDriveTypeW")
	setFileTimeProc                       = kernel32.NewProc("SetFileTime")
	procDeleteFileW                       = kernel32.NewProc("DeleteFileW")
	procCreateFile2                       = kernel32.NewProc("CreateFile2")
	procFindFirstFileNameW                = kernel32.NewProc("FindFirstFileNameW")
	procFileTimeToLocalFileTime           = kernel32.NewProc("FileTimeToLocalFileTime")
	procFindClose                         = kernel32.NewProc("FindClose")
	procFindFirstFileExW                  = kernel32.NewProc("FindFirstFileExW")
	procFindFirstFileW                    = kernel32.NewProc("FindFirstFileW")
	procFindFirstStreamW                  = kernel32.NewProc("FindFirstStreamW")
	procFindNextFileW                     = kernel32.NewProc("FindNextFileW")
	procFindNextStreamW                   = kernel32.NewProc("FindNextStreamW")
	procGetDiskFreeSpaceExW               = kernel32.NewProc("GetDiskFreeSpaceExW")
	procGetDiskFreeSpaceW                 = kernel32.NewProc("GetDiskFreeSpaceW")
	procGetFileAttributesExW              = kernel32.NewProc("GetFileAttributesExW")
	procGetFileAttributesW                = kernel32.NewProc("GetFileAttributesW")
	procGetFileSize                       = kernel32.NewProc("GetFileSize")
	procGetFileSizeEx                     = kernel32.NewProc("GetFileSizeEx")
	procGetFileTime                       = kernel32.NewProc("GetFileTime")
	procGetFileType                       = kernel32.NewProc("GetFileType")
	procGetFinalPathNameByHandleW         = kernel32.NewProc("GetFinalPathNameByHandleW")
	procGetFullPathNameW                  = kernel32.NewProc("GetFullPathNameW")
	procGetLogicalDriveStringsW           = kernel32.NewProc("GetLogicalDriveStringsW")
	procGetLongPathNameW                  = kernel32.NewProc("GetLongPathNameW")
	procReadFile                          = kernel32.NewProc("ReadFile")
	procReadFileEx                        = kernel32.NewProc("ReadFileEx")
	procSetFileTime                       = kernel32.NewProc("SetFileTime")
	procLocalFileTimeToFileTime           = kernel32.NewProc("LocalFileTimeToFileTime")
	procWriteFile                         = kernel32.NewProc("WriteFile")

	// Custom Errors for returns
	errUnknownDriveType = NewDriveError("unknown drive type")
	errNoRootDir        = NewDriveError("invalid root drive path")
)

type DriveError struct {
	err string
	msg string
}

func NewDriveError(msg string) DriveError {
	return DriveError{err:"Drive Error",msg:msg}
}

func (D DriveError) Error() string {
	return fmt.Sprintf("%s: %s", D.err,D.msg)
}
