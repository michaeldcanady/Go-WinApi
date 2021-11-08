package fileapi

import (
	"fmt"
	"syscall"
)

const (
	MAX_PATH                         = syscall.MAX_PATH
	guidBufLen                       = MAX_PATH + 1
	MaxVolumeNameLength              = 50
	MAXDWORD                         = 4294967295
	FILE_SUPPORTS_USN_JOURNAL uint32 = 0x02000000
	INVALID_HANDLE_VALUE             = HANDLE(syscall.InvalidHandle)
)

//DriveType
type DriveType int64

func (d DriveType) String() string {
	return [...]string{"DRIVE_UNKNOWN", "DRIVE_NO_ROOT_DIR", "DRIVE_REMOVABLE", "DRIVE_FIXED", "DRIVE_REMOTE", "DRIVE_CDROM", "DRIVE_RAMDISK"}[d]
}

//DRIVE_UNKNOWN The drive type cannot be determined.
const DRIVE_UNKNOWN DriveType = 0 //0
//DRIVE_NO_ROOT_DIR The root path is invalid; for example, there is no volume mounted at the specified path.
const DRIVE_NO_ROOT_DIR DriveType = 1 //1
//DRIVE_REMOVABLE The drive has removable media; for example, a floppy drive, thumb drive, or flash card reader.
const DRIVE_REMOVABLE DriveType = 2 //2
//DRIVE_FIXED The drive has fixed media; for example, a hard disk drive or flash drive.
const DRIVE_FIXED DriveType = 3 //3
//DRIVE_REMOTE The drive is a remote (network) drive.
const DRIVE_REMOTE DriveType = 4 //4
//DRIVE_CDROM The drive is a CD-ROM drive.
const DRIVE_CDROM DriveType = 5 //5
//DRIVE_RAMDISK The drive is a RAM disk.
const DRIVE_RAMDISK DriveType = 6 //6

//Possible Security Attributes

//FILE_ATTRIBUTE_READONLY The file is read only. Applications can read the file, but cannot write to or delete it.
const FILE_ATTRIBUTE_READONLY SecurityAttribute = 1

//FILE_ATTRIBUTE_HIDDEN The file is hidden. Do not include it in an ordinary directory listing.
const FILE_ATTRIBUTE_HIDDEN SecurityAttribute = 2

//FILE_ATTRIBUTE_SYSTEM The file is part of or used exclusively by an operating system.
const FILE_ATTRIBUTE_SYSTEM SecurityAttribute = 4

//FILE_ATTRIBUTE_ARCHIVE The file should be archived. Applications use this attribute to mark files for backup or removal.
const FILE_ATTRIBUTE_ARCHIVE SecurityAttribute = 32

//FILE_ATTRIBUTE_NORMAL The file does not have other attributes set. This attribute is valid only if used alone.
//
//Only Used Alone.
const FILE_ATTRIBUTE_NORMAL SecurityAttribute = 128

//FILE_ATTRIBUTE_TEMPORARY The file is being used for temporary storage.
//
//For more information, see the Caching Behavior section of this topic.
const FILE_ATTRIBUTE_TEMPORARY SecurityAttribute = 256

//FILE_ATTRIBUTE_OFFLINE The data of a file is not immediately available. This attribute indicates that file data is physically moved to offline storage.
//This attribute is used by Remote Storage, the hierarchical storage management software. Applications should not arbitrarily change this attribute.
const FILE_ATTRIBUTE_OFFLINE SecurityAttribute = 4096

//FILE_ATTRIBUTE_ENCRYPTED The file or directory is encrypted. For a file, this means that all data in the file is encrypted.
//For a directory, this means that encryption is the default for newly created files and subdirectories. For more information, see File Encryption.
//
//This flag has no effect if FILE_ATTRIBUTE_SYSTEM is also specified.
//
//This flag is not supported on Home, Home Premium, Starter, or ARM editions of Windows.
const FILE_ATTRIBUTE_ENCRYPTED SecurityAttribute = 16384

//AdditionalFlags
type AdditionalFlags int64

func (a AdditionalFlags) String() string {
	return [...]string{"", "FIND_FIRST_EX_CASE_SENSITIVE", "FIND_FIRST_EX_LARGE_FETCH", "", "FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY"}[a]
}

//FIND_FIRST_EX_CASE_SENSITIVE Searches are case-sensitive.
const FIND_FIRST_EX_CASE_SENSITIVE AdditionalFlags = 1

//FIND_FIRST_EX_LARGE_FETCH Uses a larger buffer for directory queries, which can increase performance of the find operation.
//Windows Server 2008, Windows Vista, Windows Server 2003 and Windows XP:  This value is not supported until Windows Server 2008 R2 and Windows 7.
const FIND_FIRST_EX_LARGE_FETCH AdditionalFlags = 2

//FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY Limits the results to files that are physically on disk. This flag is only relevant when a file virtualization filter is present.
const FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY AdditionalFlags = 4

type InfoLevel int64

func (i InfoLevel) String() string {
	return [...]string{"FindExInfoStandard", "FindExInfoBasic", "FindExInfoMaxInfoLevel"}[i]
}

//FINDEX_INFO_LEVELS
const FINDEX_INFO_STANDARD InfoLevel = 0
const FINDEX_INFO_BASIC InfoLevel = 1
const FINDEX_INFO_MAX_INFO_LEVEL InfoLevel = 2

//SearchOps
type SearchOps int64

//FINDEX_SEARCH_NAME_MATCH The search for a file that matches a specified file name.
//
//The lpSearchFilter parameter of
//FindFirstFileEx must be
//NULL when this search operation is used.
const FINDEX_SEARCH_NAME_MATCH SearchOps = 1

//FINDEX_SEARCH_LIMIT_TO_DIRECTORIES This is an advisory flag. If the file system supports directory filtering, the function searches for a file that matches the specified name
//and is also a directory. If the file system does not support directory filtering, this flag is silently ignored.
//
//The lpSearchFilter parameter of the FindFirstFileEx function must be NULL when this search value is used.
//If directory filtering is desired, this flag can be used on all file systems, but because it is an advisory flag and only affects file systems that support it,
//the application must examine the file attribute data stored in the lpFindFileData parameter of the FindFirstFileEx function to determine whether the function has returned a
//handle to a directory.
const FINDEX_SEARCH_LIMIT_TO_DIRECTORIES SearchOps = 2

//FINDEX_SEARCH_LIMIT_TO_DEVICES This filtering type is not available.
//
//For more information, see
//Device Interface Classes.
const FINDEX_SEARCH_LIMIT_TO_DEVICES SearchOps = 3
const FINDEX_SEARCH_MAX_SEARCH_OP SearchOps = 4

//FileType
type FileType int64

//FILE_TYPE_UNKNOWN Either the type of the specified file is unknown, or the function failed.
const FILE_TYPE_UNKNOWN FileType = 0x0000

//FILE_TYPE_DISK The specified file is a disk file.
const FILE_TYPE_DISK FileType = 0x0001

//FILE_TYPE_CHAR The specified file is a character file, typically an LPT device or a console.
const FILE_TYPE_CHAR FileType = 0x0002

//FILE_TYPE_PIPE The specified file is a socket, a named pipe, or an anonymous pipe.
const FILE_TYPE_PIPE FileType = 0x0003

//FILE_TYPE_REMOTE Unused.
const FILE_TYPE_REMOTE FileType = 0x8000

type SecurityAttributeError struct {
	err string
	msg string
}

func NewSecurityAttributeError(msg string) SecurityAttributeError {
	return SecurityAttributeError{err: "Security Attribute Error:", msg: msg}
}

func (S SecurityAttributeError) Error() string {
	return fmt.Sprintf("%s: %s", S.err, S.msg)
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
	return fmt.Sprintf("%s: %s", L.err, L.msg)
}
