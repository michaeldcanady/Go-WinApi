package fileapi

import (
	"syscall"
)

type (
	//DriveType
	DriveType int64

	//AdditionalFlags
	AdditionalFlags int64

	//InfoLevel
	InfoLevel int64

	//SearchOps
	SearchOps int64

	//FileType
	FileType int64

	//CreationDisposition
	CreationDisposition int64

	//ShareMode The requested sharing mode of the file or device, which can be read, write, both, delete, all of these, or none (refer to the following table). Access requests to attributes or extended attributes are not affected
	//by this flag.
	//
	// If this parameter is zero and CreateFile succeeds, the file or device cannot be shared and cannot be opened again until the handle to the file or device is closed. For more information, see the Remarks section.
	//
	//You cannot request a sharing mode that conflicts with the access mode that is specified in an existing request that has an open handle. CreateFile would fail and the GetLastError function would return
	//ERROR_SHARING_VIOLATION.
	//
	//To enable a process to share a file or device while another process has the file or device open, use a compatible combination of one or more of the following values.
	//For more information about valid combinations of this parameter with the dwDesiredAccess parameter, see Creating and Opening Files.
	ShareMode int64

	SecurityAttribute int64
)

func (d DriveType) String() string {
	return [...]string{"DRIVE_UNKNOWN", "DRIVE_NO_ROOT_DIR", "DRIVE_REMOVABLE", "DRIVE_FIXED", "DRIVE_REMOTE", "DRIVE_CDROM", "DRIVE_RAMDISK"}[d]
}

func (a AdditionalFlags) String() string {
	return [...]string{"", "FIND_FIRST_EX_CASE_SENSITIVE", "FIND_FIRST_EX_LARGE_FETCH", "", "FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY"}[a]
}

func (i InfoLevel) String() string {
	return [...]string{"FindExInfoStandard", "FindExInfoBasic", "FindExInfoMaxInfoLevel"}[i]
}

//String
func (c CreationDisposition) String() string {
	return [...]string{"CREATE_NEW", "CREATE_ALWAYS", "OPEN_EXISTING", "OPEN_ALWAYS", "TRUNCATE_EXISTING"}[c]
}

const (
	MAX_PATH            = 260
	guidBufLen          = MAX_PATH + 1
	MaxVolumeNameLength = 50
	MAXDWORD            = 4294967295
	NUL                 = 0x0000
	errorMoreData       = 234

	INVALID_HANDLE_VALUE HANDLE = 18446744073709551615
	//DRIVE_UNKNOWN The drive type cannot be determined.
	DRIVE_UNKNOWN DriveType = 0

	//DRIVE_NO_ROOT_DIR The root path is invalid; for example, there is no volume mounted at the specified path.
	DRIVE_NO_ROOT_DIR DriveType = 1

	//DRIVE_REMOVABLE The drive has removable media; for example, a floppy drive, thumb drive, or flash card reader.
	DRIVE_REMOVABLE DriveType = 2

	//DRIVE_FIXED The drive has fixed media; for example, a hard disk drive or flash drive.
	DRIVE_FIXED DriveType = 3

	//DRIVE_REMOTE The drive is a remote (network) drive.
	DRIVE_REMOTE DriveType = 4

	//DRIVE_CDROM The drive is a CD-ROM drive.
	DRIVE_CDROM DriveType = 5

	//DRIVE_RAMDISK The drive is a RAM disk.
	DRIVE_RAMDISK DriveType = 6
	//Possible Security Attributes

	//FILE_ATTRIBUTE_READONLY The file is read only. Applications can read the file, but cannot write to or delete it.
	FILE_ATTRIBUTE_READONLY SecurityAttribute = 1

	//FILE_ATTRIBUTE_HIDDEN The file is hidden. Do not include it in an ordinary directory listing.
	FILE_ATTRIBUTE_HIDDEN SecurityAttribute = 2

	//FILE_ATTRIBUTE_SYSTEM The file is part of or used exclusively by an operating system.
	FILE_ATTRIBUTE_SYSTEM SecurityAttribute = 4

	//FILE_ATTRIBUTE_ARCHIVE The file should be archived. Applications use this attribute to mark files for backup or removal.

	FILE_ATTRIBUTE_ARCHIVE SecurityAttribute = 32

	//FILE_ATTRIBUTE_NORMAL The file does not have other attributes set. This attribute is valid only if used alone.
	//
	//Only Used Alone.
	FILE_ATTRIBUTE_NORMAL SecurityAttribute = 128

	//FILE_ATTRIBUTE_TEMPORARY The file is being used for temporary storage.
	//
	//For more information, see the Caching Behavior section of this topic.
	FILE_ATTRIBUTE_TEMPORARY SecurityAttribute = 256

	//FILE_ATTRIBUTE_OFFLINE The data of a file is not immediately available. This attribute indicates that file data is physically moved to offline storage.
	//This attribute is used by Remote Storage, the hierarchical storage management software. Applications should not arbitrarily change this attribute.
	FILE_ATTRIBUTE_OFFLINE SecurityAttribute = 4096

	//FILE_ATTRIBUTE_ENCRYPTED The file or directory is encrypted. For a file, this means that all data in the file is encrypted.
	//For a directory, this means that encryption is the default for newly created files and subdirectories. For more information, see File Encryption.
	//
	//This flag has no effect if FILE_ATTRIBUTE_SYSTEM is also specified.
	//
	//This flag is not supported on Home, Home Premium, Starter, or ARM editions of Windows.
	FILE_ATTRIBUTE_ENCRYPTED SecurityAttribute = 16384

	//FIND_FIRST_EX_CASE_SENSITIVE Searches are case-sensitive.
	FIND_FIRST_EX_CASE_SENSITIVE AdditionalFlags = 1

	//FIND_FIRST_EX_LARGE_FETCH Uses a larger buffer for directory queries, which can increase performance of the find operation.
	//Windows Server 2008, Windows Vista, Windows Server 2003 and Windows XP:  This value is not supported until Windows Server 2008 R2 and Windows 7.
	FIND_FIRST_EX_LARGE_FETCH AdditionalFlags = 2

	//FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY Limits the results to files that are physically on disk. This flag is only relevant when a file virtualization filter is present.
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY AdditionalFlags = 4
	//FINDEX_INFO_STANDARD The FindFirstFileEx function retrieves a
	//standard set of attribute information. The data is returned in a
	//WIN32_FIND_DATA structure.
	FINDEX_INFO_STANDARD InfoLevel = 0
	//FINDEX_INFO_BASIC The FindFirstFileEx function does not query the short file name, improving overall enumeration speed. The data is returned in a
	//WIN32_FIND_DATA structure, and the cAlternateFileName
	//member is always a NULL string.
	FINDEX_INFO_BASIC InfoLevel = 1
	//FINDEX_INFO_MAX_INFO_LEVEL This value is used for validation. Supported values are less than this value.
	FINDEX_INFO_MAX_INFO_LEVEL InfoLevel = 2
	//FINDEX_SEARCH_NAME_MATCH The search for a file that matches a specified file name.
	//
	//The lpSearchFilter parameter of
	//FindFirstFileEx must be
	//NULL when this search operation is used.
	FINDEX_SEARCH_NAME_MATCH SearchOps = 1

	//FINDEX_SEARCH_LIMIT_TO_DIRECTORIES This is an advisory flag. If the file system supports directory filtering, the function searches for a file that matches the specified name
	//and is also a directory. If the file system does not support directory filtering, this flag is silently ignored.
	//
	//The lpSearchFilter parameter of the FindFirstFileEx function must be NULL when this search value is used.
	//If directory filtering is desired, this flag can be used on all file systems, but because it is an advisory flag and only affects file systems that support it,
	//the application must examine the file attribute data stored in the lpFindFileData parameter of the FindFirstFileEx function to determine whether the function has returned a
	//handle to a directory.
	FINDEX_SEARCH_LIMIT_TO_DIRECTORIES SearchOps = 2

	//FINDEX_SEARCH_LIMIT_TO_DEVICES This filtering type is not available.
	//
	//For more information, see
	//Device Interface Classes.
	FINDEX_SEARCH_LIMIT_TO_DEVICES SearchOps = 3

	//FINDEX_SEARCH_MAX_SEARCH_OP Undefined by documentation.
	FINDEX_SEARCH_MAX_SEARCH_OP SearchOps = 4

	//FILE_TYPE_UNKNOWN Either the type of the specified file is unknown, or the function failed.
	FILE_TYPE_UNKNOWN FileType = 0x0000

	//FILE_TYPE_DISK The specified file is a disk file.
	FILE_TYPE_DISK FileType = 0x0001

	//FILE_TYPE_CHAR The specified file is a character file, typically an LPT device or a console.
	FILE_TYPE_CHAR FileType = 0x0002

	//FILE_TYPE_PIPE The specified file is a socket, a named pipe, or an anonymous pipe.
	FILE_TYPE_PIPE FileType = 0x0003

	//FILE_TYPE_REMOTE Unused.
	FILE_TYPE_REMOTE FileType = 0x8000

	//CREATE_NEW Creates a new file, only if it does not already exist.
	//
	//If the specified file exists, the function fails and the last-error code is set to ERROR_FILE_EXISTS (80).
	//
	//If the specified file does not exist and is a valid path to a writable location, a new file is created.
	CREATE_NEW CreationDisposition = 1

	//CREATE_ALWAYS Creates a new file, always.
	//
	//If the specified file exists and is writable, the function overwrites the file, the function succeeds, and last-error code is set to ERROR_ALREADY_EXISTS (183).
	//
	//If the specified file does not exist and is a valid path, a new file is created, the function succeeds, and the last-error code is set to zero.
	//
	//For more information, see the Remarks section of this topic.
	CREATE_ALWAYS CreationDisposition = 2

	//OPEN_EXISTING Opens a file or device, only if it exists.
	//
	//If the specified file or device does not exist, the function fails and the last-error code is set to ERROR_FILE_NOT_FOUND (2).
	//
	//For more information about devices, see the Remarks section.
	OPEN_EXISTING CreationDisposition = 3

	//OPEN_ALWAYS Opens a file, always.
	//
	//If the specified file exists, the function succeeds and the last-error code is set to ERROR_ALREADY_EXISTS (183).
	//
	//If the specified file does not exist and is a valid path to a writable location, the function creates a file and the last-error code is set to zero.
	OPEN_ALWAYS CreationDisposition = 4

	//TRUNCATE_EXISTING Opens a file and truncates it so that its size is zero bytes, only if it exists.
	//If the specified file does not exist, the function fails and the last-error code is set to ERROR_FILE_NOT_FOUND (2).
	//
	//The calling process must open the file with the GENERIC_WRITE bit set as part of the dwDesiredAccess parameter.
	TRUNCATE_EXISTING CreationDisposition = 5

	//DO_NOT_FILE_SHARE Prevents other processes from opening a file or device if they request delete, read, or write access.
	DO_NOT_FILE_SHARE = 0x00000000

	//FILE_SHARE_READ Enables subsequent open operations on a file or device to request delete access.
	//
	//Otherwise, other processes cannot open the file or device if they request delete access.
	//
	//If this flag is not specified, but the file or device has been opened for delete access, the function fails.
	//
	//Note  Delete access allows both delete and rename operations.
	FILE_SHARE_READ = 0x00000001

	//FILE_SHARE_WRITE Enables subsequent open operations on a file or device to request read access.
	//Otherwise, other processes cannot open the file or device if they request read access.
	//
	//If this flag is not specified, but the file or device has been opened for read access, the function fails.
	FILE_SHARE_WRITE = 0x00000002

	//FILE_SHARE_DELETE Enables subsequent open operations on a file or device to request write access. Otherwise, other processes cannot open the file or device if they request write access.
	//If this flag is not specified, but the file or device has been opened for write access or has a file mapping with write access, the function fails.
	FILE_SHARE_DELETE = 0x00000004

	//GENERIC_READ
	GENERIC_READ = 2147483648

	//FILE_FLAG_BACKUP_SEMANTICS
	FILE_FLAG_BACKUP_SEMANTICS = 2147483648
)

type HANDLE syscall.Handle

func (H HANDLE) ToUintPtr() uintptr {
	return uintptr(H)
}

type SecurityAttributeError struct {
	err string
	msg string
}

func NewSecurityAttributeError(msg string) SecurityAttributeError {
	return SecurityAttributeError{err: "Security Attribute Error:", msg: msg}
}

func (S SecurityAttributeError) Error() string {
	return S.err + ": " + S.msg
}

//LogicalDriveError Error given when there is a problem in a function related to logical drives
type LogicalDriveError struct {
	err string
	msg string
}

//NewLogicalDriveError Creates a new LogicalDriveError
func NewLogicalDriveError(msg string) LogicalDriveError {
	return LogicalDriveError{err: "Logical Drive Error:", msg: msg}
}

//Error returns the LogicalDriveError
func (L LogicalDriveError) Error() string {
	return L.err + ": " + L.msg
}

type DriveError struct {
	err string
	msg string
}

func NewDriveError(msg string) DriveError {
	return DriveError{err: "Drive Error", msg: msg}
}

func (D DriveError) Error() string {
	return D.err + ": " + D.msg
}

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procCreateFileW                       = kernel32.NewProc("CreateFileW")
	procCreateDirectoryW                  = kernel32.NewProc("CreateDirectoryW")
	areFileApisANSIProc                   = kernel32.NewProc("AreFileApisANSI")
	findFirstVolumeWProc                  = kernel32.NewProc("FindFirstVolumeW")
	findNextVolumeWProc                   = kernel32.NewProc("FindNextVolumeW")
	procFindVolumeClose                   = kernel32.NewProc("FindVolumeClose")
	getVolumePathNamesForVolumeNameWProc  = kernel32.NewProc("GetVolumePathNamesForVolumeNameW")
	procGetVolumeNameForVolumeMountPointW = kernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procGetVolumeInformationW             = kernel32.NewProc("GetVolumeInformationW")
	procGetLogicalDrives                  = kernel32.NewProc("GetLogicalDrives")
	procDeleteVolumeMountPointW           = kernel32.NewProc("DeleteVolumeMountPointW")
	procGetDriveTypeW                     = kernel32.NewProc("GetDriveTypeW")
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
