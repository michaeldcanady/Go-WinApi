package GoWinApi

import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"
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

	fileType = map[int64]string{
		0x0001: "FILE_TYPE_DISK",
		0x0002: "FILE_TYPE_CHAR",
		0x0003: "FILE_TYPE_PIPE",
		0x8000: "FILE_TYPE_REMOTE",
	}

	FileNameFlags = map[int64]string{
		0x0: "FILE_NAME_NORMALIZED",
		0x8: "FILE_NAME_OPENED",
	}
)

//AreFileApisANSI If the set of file I/O functions is using the ANSI code page, the return value is true.
//If the set of file I/O functions is using the OEM code page, the return value is false.
func AreFileApisANSI() bool {
	ret, _, _ := areFileApisANSIProc.Call()

	return ret != 0
}

//Creates a new directory. If the underlying file system supports security on files and directories, the function applies a specified
//security descriptor to the new directory.
//
// To specify a template directory, use the CreateDirectoryEx function.
//
// To perform this operation as a transacted operation, use the CreateDirectoryTransacted function.

func CreateDirectoryW(pathName string, SecurityAttributes SecurityAttribute) error {

	r, _, err := procCreateDirectoryW.Call(
		UintptrFromString(pathName),
		uintptr(SecurityAttributes),
	)

	if r == 0 {
		return err
	}

	return nil
}

//CreateFile2 Creates or opens a file or I/O device. The most commonly used I/O devices are as follows: file, file stream, directory, physical disk, volume,
//console buffer, tape drive, communications resource, mailslot, and pipe. The function returns a handle that can be used to access the file or device for various types
//of I/O depending on the file or device and the flags and attributes specified.
//
//When called from a Windows Store app, CreateFile2 is simplified. You can open only files or directories inside the ApplicationData.LocalFolder or Package.
//InstalledLocation directories. You can't open named pipes or mailslots or create encrypted files (FILE_ATTRIBUTE_ENCRYPTED).
func CreateFile2(fileName string, dwDesiredAccess, dwShareMode, dwCreationDisposition uint32) (handle HANDLE, err error) {

	ret, _, err := procCreateFile2.Call(
		UintptrFromString(fileName),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(dwCreationDisposition),
		0,
	)

	handle = HANDLE(ret)

	if handle == INVALID_HANDLE_VALUE {
		return handle, err
	}

	return handle, nil
}

//CreateFileW Creates or opens a file or I/O device. The most commonly used I/O devices are as follows: file, file stream, directory, physical disk, volume, console buffer, tape drive, communications resource, mailslot, and pipe. The function returns a handle that can be used to access the file or device for various types of I/O depending on the file or device and the flags and attributes specified.
//To perform this operation as a transacted operation, which results in a handle that can be used for transacted I/O, use the CreateFileTransacted function.
func CreateFileW(fileName string, dwDesiredAccess, dwShareMode, lpSecurityAttributes SecurityAttribute, dwCreationDisposition, dwFlagsAndAttributes uint32, dhTemplateFile HANDLE) (HANDLE, error) {

	ret, _, err := procCreateFileW.Call(
		UintptrFromString(fileName),    // [in] LPCTSTR
		uintptr(dwDesiredAccess),       // [in] DWORD
		uintptr(dwShareMode),           // [in] DWORD
		uintptr(lpSecurityAttributes),  // [in] LPSECURITY_ATT...
		uintptr(dwCreationDisposition), // [in] DWORD
		uintptr(dwFlagsAndAttributes),  // [in] DWORD
		uintptr(dhTemplateFile),        // [in] HANDLE
	)

	if ret == 0 || ret == INVALID_HANDLE_VALUE.ToUintPtr() {
		return HANDLE(0), err
	}

	return HANDLE(ret), nil
}

//DeleteFileW Deletes an existing file.
//
//To perform this operation as a transacted operation, use the DeleteFileTransacted function.
func DeleteFileW(fileName string) error {

	ret, _, err := procDeleteFileW.Call(UintptrFromString(fileName))

	if ret == 0 {
		return err
	}

	return nil
}

//DeleteVolumeMountPointW Deletes a drive letter or mounted folder.
func DeleteVolumeMountPointW(volumeMountPoint string) error {

	ret, _, err := procDeleteVolumeMountPointW.Call(UintptrFromString(volumeMountPoint))

	if ret == 0 {
		return err
	}

	return nil
}

//FileTimeToLocalFileTime Converts a file time to a local file time.
func FileTimeToLocalFileTime(lpFileTime FILETIME) (lpLocalFileTime FILETIME, err error) {

	ret, _, err := procFileTimeToLocalFileTime.Call(
		lpFileTime.ToUintPtr(),
		lpLocalFileTime.ToUintPtr(),
	)
	if ret == 0 {
		return lpLocalFileTime, err
	}

	return lpLocalFileTime, nil
}

//FindClose Closes a file search handle opened by the FindFirstFile, FindFirstFileEx, FindFirstFileNameW, FindFirstFileNameTransactedW, FindFirstFileTransacted,
//FindFirstStreamTransactedW, or FindFirstStreamW functions.
func FindClose(hFindFile HANDLE) error {

	ret, _, err := procFindClose.Call(hFindFile.ToUintPtr())

	if ret == 0 {
		return err
	}

	return nil
}

//FindFirstFileExW Searches a directory for a file or subdirectory with a name and attributes that match those specified.
//For the most basic version of this function, see FindFirstFile.
//To perform this operation as a transacted operation, use the FindFirstFileTransacted function.
func FindFirstFileExW(FileName string, fInfoLevelId int32, fSearchOp int32, dwAdditionalFlags AdditionalFlags) (HANDLE, Win32FindDataW, error) {

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

//TODO Figure out how to get string length to work
func FindFirstFileNameW(fileName string) {

	var LinkName uintptr
	var stringLength uint32

	ret, _, err := procFindFirstFileNameW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&stringLength)),
		LinkName,
	)

	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println(LinkName)
}

//FindFirstFileW Searches a directory for a file or subdirectory with a name that matches a specific name (or partial name if wildcards are used).
//
//To specify additional attributes to use in a search, use the FindFirstFileEx function.
//
//To perform this operation as a transacted operation, use the FindFirstFileTransacted function.
func FindFirstFileW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstFileW.Call(
		UintptrFromString(fileName),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 || HANDLE(ret) == INVALID_HANDLE_VALUE {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}

//FindFirstStreamW Enumerates the first stream with a ::$DATA stream type in the specified file or directory.
//
//To perform this operation as a transacted operation, use the FindFirstStreamTransactedW function.
func FindFirstStreamW(fileName string) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindFirstStreamW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&lpFindFileData)),
		0,
	)

	if ret == 0 {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}

//FindFirstVolume Retrieves the name of a volume on a computer. FindFirstVolume is used to begin scanning the volumes of a computer.
func FindFirstVolume() (HANDLE, string, error) {
	const invalidHandleValue = ^uintptr(0)

	guid := make([]uint16, guidBufLen)

	handle, _, err := findFirstVolumeWProc.Call(
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if handle == invalidHandleValue {
		return INVALID_HANDLE_VALUE, "", err
	}

	return HANDLE(handle), syscall.UTF16ToString(guid), nil
}

//FindNextFileW Continues a file search from a previous call to the FindFirstFile, FindFirstFileEx, or FindFirstFileTransacted functions.
func FindNextFileW(hFindFile HANDLE) (Win32FindDataW, error) {

	//Imediately returns an error if the handle passed in is invalid
	if hFindFile == HANDLE(syscall.InvalidHandle) {
		return Win32FindDataW{}, errors.New("Invalid handle passed")
	}

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextFileW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 {
		return Win32FindDataW{}, err
	}

	return newWin32FindData(lpFindFileData), nil
}

//FindNextStreamW Continues a stream search started by a previous call to the FindFirstStreamW function.
func FindNextStreamW(hFindFile HANDLE) (HANDLE, Win32FindDataW, error) {

	var lpFindFileData WIN32_FIND_DATAA

	ret, _, err := procFindNextStreamW.Call(
		uintptr(hFindFile),
		uintptr(unsafe.Pointer(&lpFindFileData)),
	)

	if ret == 0 {
		return HANDLE(0), Win32FindDataW{}, err
	}

	return HANDLE(ret), newWin32FindData(lpFindFileData), nil
}

//FindNextVolume Continues a volume search started by a call to the FindFirstVolume function. FindNextVolume finds one volume per call.
func FindNextVolume(handle HANDLE) (string, bool, error) {
	const noMoreFiles = 18

	guid := make([]uint16, guidBufLen)

	rc, _, err := findNextVolumeWProc.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if rc == 1 {
		return syscall.UTF16ToString(guid), true, nil
	}

	if err.(syscall.Errno) == noMoreFiles {
		return "", false, nil
	}
	return "", false, err
}

//FindVolumeClose Closes the specified volume search handle. The FindFirstVolume and FindNextVolume functions use this search handle to locate volumes.
func FindVolumeClose(handle HANDLE) error {
	ok, _, err := procFindVolumeClose.Call(uintptr(handle))

	if ok == 0 {
		return err
	}

	return nil
}

//GetDiskFreeSpaceExW Retrieves information about the amount of space that is available on a disk volume, which is the total amount of space, the total amount of
//free space, and the total amount of free space available to the user that is associated with the calling thread.
func GetDiskFreeSpaceExW(directoryName string) (int64, int64, int64, error) {

	var lpFreeBytesAvailableToCaller, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes int64

	ret, _, err := procGetDiskFreeSpaceExW.Call(
		UintptrFromString(directoryName),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailableToCaller)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)),
	)
	if ret == 0 {
		return 0, 0, 0, err
	}

	return lpFreeBytesAvailableToCaller, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes, nil
}

//GetDiskFreeSpaceW Retrieves information about the specified disk, including the amount of free space on the disk.
func GetDiskFreeSpaceW(lpDirectoryName string) (lpSectorsPerCluster uint32, lpBytesPerSector uint32, lpNumberOfFreeClusters uint32, lpTotalNumberOfClusters uint32, err error) {

	ret, _, err := procGetDiskFreeSpaceW.Call(
		UintptrFromString(lpDirectoryName),
		uintptr(unsafe.Pointer(&lpSectorsPerCluster)),
		uintptr(unsafe.Pointer(&lpBytesPerSector)),
		uintptr(unsafe.Pointer(&lpNumberOfFreeClusters)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfClusters)),
	)

	if ret == 0 {
		return 0, 0, 0, 0, err
	}

	return lpSectorsPerCluster, lpBytesPerSector, lpNumberOfFreeClusters, lpTotalNumberOfClusters, nil
}

//GetDriveTypeW Determines whether a disk drive is a removable, fixed, CD-ROM, RAM disk, or network drive.
//
//To determine whether a drive is a USB-type drive, call SetupDiGetDeviceRegistryProperty and specify the SPDRP_REMOVAL_POLICY property.
func GetDriveTypeW(PathName string) (string, error) {

	ret, _, err := procGetDriveTypeW.Call(UintptrFromString(PathName))

	return DriveType(ret).String(), err
}

//GetFileAttributesExW Retrieves attributes for a specified file or directory.
//
//To perform this operation as a transacted operation, use the GetFileAttributesTransacted function.
func GetFileAttributesExW(fileName string) (Win32FileAttributeData, error) {

	var lpFileInformation Win32FileAttributeDataA

	ret, _, err := procGetFileAttributesExW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&lpFileInformation)),
	)

	if ret == 0 {
		return newWin32FileAttributeData(lpFileInformation), err
	}

	return newWin32FileAttributeData(lpFileInformation), nil
}

//GetFileAttributesW Retrieves file system attributes for a specified file or directory.
//
//To get more attribute information, use the GetFileAttributesEx function.
//
//To perform this operation as a transacted operation, use the GetFileAttributesTransacted function.
func GetFileAttributesW(fileName string) ([]string, error) {
	ret, _, err := procGetFileAttributesW.Call(UintptrFromString(fileName))

	if ret == 0xFFFFFFFF {
		return []string{}, err
	}

	return SeperateFlags(uint32(ret), dwFileAttributeFlags), err
}

//GetFileSize Retrieves the size of the specified file, in bytes.
//
//It is recommended that you use GetFileSizeEx.
func GetFileSize(hFile HANDLE) (int64, error) {

	ret, _, err := procGetFileSize.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0xFFFFFFFF {
		return 0, err
	}

	return int64(ret), nil
}

//GetFileSizeEx Retrieves the size of the specified file.
func GetFileSizeEx(hFile HANDLE) (int64, error) {

	ret, _, err := procGetFileSizeEx.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0 {
		return 0, err
	}

	return int64(ret), nil
}

//GetFileTime Retrieves the date and time that a file or directory was created, last accessed, and last modified.
func GetFileTime(hFile HANDLE) (time.Time, time.Time, time.Time, error) {

	var dwCreationTime, dwLastAccessTime, dwLastWriteTime FILETIME

	ret, _, err := procGetFileTime.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&dwCreationTime)),
		uintptr(unsafe.Pointer(&dwLastAccessTime)),
		uintptr(unsafe.Pointer(&dwLastWriteTime)),
	)

	if ret == 0 {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	CreationTime, err := FileTimeToSystemTime(dwCreationTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}
	LastAccessTime, err := FileTimeToSystemTime(dwLastAccessTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}
	LastWriteTime, err := FileTimeToSystemTime(dwLastWriteTime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	return CreationTime, LastAccessTime, LastWriteTime, nil
}

//GetFileType Retrieves the file type of the specified file.
func GetFileType(hFile HANDLE) (string, error) {

	ret, _, err := procGetFileType.Call(uintptr(hFile))

	if ret == 0 {
		return "FILE_TYPE_UNKNOWN", err
	}

	return SeperateFlags(uint32(ret), fileType)[0], nil
}

//GetFinalPathNameByHandleW Retrieves the final path for the specified file.
//
//For more information about file and path names, see Naming a File.
func GetFinalPathNameByHandleW(hFile HANDLE) (string, string, error) {

	var bufSize uint32 = syscall.MAX_PATH // 260
	buf := make([]uint16, bufSize)
	var rawFlags uint32

	ret, _, err := procGetFinalPathNameByHandleW.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&rawFlags)),
	)

	if ret == 0 {
		return "", "", err
	}

	return syscall.UTF16ToString(buf), SeperateFlags(rawFlags, FileNameFlags)[0], nil
}

//GetFullPathNameW Retrieves the full path and file name of the specified file.
//
//To perform this operation as a transacted operation, use the GetFullPathNameTransacted function.
//
//For more information about file and path names, see File Names, Paths, and Namespaces.
func GetFullPathNameW(lpFileName string) (string, error) {

	var bufSize uint32 = syscall.MAX_PATH // 260
	buf := make([]uint16, bufSize)

	ret, _, err := procGetFullPathNameW.Call(
		UintptrFromString(lpFileName),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&buf[0])),
		0,
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil
}

//GetLogicalDrives returns a list of all logical drives on the host machine
func GetLogicalDrives() ([]string, error) {
	ret, _, _ := procGetLogicalDrives.Call()
	if ret == 0 {
		return []string{}, NewLogicalDriveError("No Drives Found")
	}
	return bitsToDrives(uint32(ret)), nil
}

//GetLogicalDriveStringsW Fills a buffer with strings that specify valid drives in the system.
func GetLogicalDriveStringsW() (string, error) {

	buf := make([]uint16, MAX_PATH)

	ret, _, err := procGetLogicalDriveStringsW.Call(
		uintptr(MAX_PATH),
		uintptr(unsafe.Pointer(&buf[0])),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil
}

//GetLongPathNameW Converts the specified path to its long form.
//
//To perform this operation as a transacted operation, use the GetLongPathNameTransacted function.
//
//For more information about file and path names, see Naming Files, Paths, and Namespaces.
func GetLongPathNameW(shortPath string) (string, error) {

	lpszLongPath := make([]uint16, MAX_PATH)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(shortPath),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(MAX_PATH),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}

//GetShortPathNameW Retrieves the short path form of the specified path.
//
//For more information about file and path names, see Naming Files, Paths, and Namespaces.
func GetShortPathNameW(longPath string) (string, error) {

	lpszShortPath := make([]uint16, MAX_PATH)

	ret, _, err := procGetLongPathNameW.Call(
		UintptrFromString(longPath),
		uintptr(unsafe.Pointer(&lpszShortPath[0])),
		uintptr(MAX_PATH),
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszShortPath), nil
}

//GetVolumeInformationW Retrieves information about the file system and volume associated with the specified root directory.
//
//To specify a handle when retrieving this information, use the GetVolumeInformationByHandleW function.
//
//To retrieve the current compression state of a file or directory, use FSCTL_GET_COMPRESSION.
func GetVolumeInformationW(rootPathName string) (volume Volume, err error) {
	if !strings.HasSuffix(rootPathName, "\\") {
		rootPathName = rootPathName + "\\"
	}

	var (
		VolumeSerialNumber     uint32
		MaximumComponentLength uint32
		FileSystemFlags        uint32
		VolumeNameBuffer              = make([]uint16, syscall.MAX_PATH+1)
		nVolumeNameSize               = uint32(len(VolumeNameBuffer))
		FileSystemNameBuffer          = make([]uint16, 255)
		nFileSystemNameSize    uint32 = syscall.MAX_PATH + 1
	)

	ret, _, err := procGetVolumeInformationW.Call(
		UintptrFromString(rootPathName),
		uintptr(unsafe.Pointer(&VolumeNameBuffer[0])),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(&VolumeSerialNumber)),
		uintptr(unsafe.Pointer(&MaximumComponentLength)),
		uintptr(unsafe.Pointer(&FileSystemFlags)),
		uintptr(unsafe.Pointer(&FileSystemNameBuffer[0])),
		uintptr(nFileSystemNameSize),
		0)
	// If GetVolumeInformationW Call returns an error
	if ret == 0 {
		return
	}

	volume = newVolume(rootPathName, VolumeNameBuffer, nVolumeNameSize, VolumeSerialNumber, MaximumComponentLength, FileSystemFlags, nFileSystemNameSize, FileSystemNameBuffer)
	//Returns new volume
	return volume, nil
}

//GetVolumeNameForVolumeMountPointW Retrieves a volume GUID path for the volume that is associated with the specified volume mount point ( drive letter, volume
//GUID path, or mounted folder).
func GetVolumeNameForVolumeMountPointW(volumeMountPoint string) (string, error) {

	if len(volumeMountPoint) == 0 {
		return "", syscall.EINVAL
	}
	if !strings.HasSuffix(volumeMountPoint, "\\") {
		volumeMountPoint = volumeMountPoint + "\\"
	}

	var vnBuffer [MaxVolumeNameLength]uint16
	p0 := &vnBuffer[0]

	re, _, err := procGetVolumeNameForVolumeMountPointW.Call(
		UintptrFromString(volumeMountPoint),
		uintptr(unsafe.Pointer(p0)),
		uintptr(MaxVolumeNameLength),
	)
	if re == 0 {
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(vnBuffer[:]), nil
}

//GetVolumePathNamesForVolumeNameW Retrieves a list of drive letters and mounted folder paths for the specified volume.
func GetVolumePathNamesForVolumeNameW(volName []string) ([]string, error) {

	var (
		pathNamesLen uint32
		pathNames    []uint16
	)

	pathNamesLen = 2
	for {
		pathNames = make([]uint16, pathNamesLen)
		pathNamesLen *= 2

		rc, _, err := getVolumePathNamesForVolumeNameWProc.Call(
			UintptrFromString(volName[0]),
			uintptr(unsafe.Pointer(&pathNames[0])),
			uintptr(pathNamesLen),
			uintptr(unsafe.Pointer(&pathNamesLen)),
		)

		if rc == 0 {
			if err.(syscall.Errno) == errorMoreData {
				continue
			}

			return nil, err
		}

		pathNames = pathNames[:pathNamesLen]
		break
	}

	var out [][]uint16
	i := 0
	for j, c := range pathNames {
		if c == NUL && i < j {
			out = append(out, pathNames[i:j+1])
			i = j + 1
		}
	}

	return LPSTRsToStrings(out), nil
}

//LocalFileTimeToFileTime Converts a local file time to a file time based on the Coordinated Universal Time (UTC).
func LocalFileTimeToFileTime(in FILETIME) (out FILETIME, err error) {

	ret, _, err := procLocalFileTimeToFileTime.Call(
		in.ToUintPtr(),
		out.ToUintPtr(),
	)
	if ret == 0 {
		return out, err
	}
	return out, nil
}

//ReadFile Reads data from the specified file or input/output (I/O) device. Reads occur at the position specified by the file pointer if supported by the device.
//
//This function is designed for both synchronous and asynchronous operations. For a similar function designed solely for asynchronous operation, see ReadFileEx.
func ReadFile(hFile HANDLE) (string, error) {

	var bufSize, err = GetFileSize(hFile)

	if err != nil {
		panic(err)
	}

	lpszLongPath := make([]uint16, bufSize)

	//var ran uint32

	ret, _, err := procReadFile.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&bufSize)),
		0,
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}

//ReadFileEx Reads data from the specified file or input/output (I/O) device. It reports its completion status asynchronously, calling the specified completion
//routine when reading is completed or canceled and the calling thread is in an alertable wait state.
func ReadFileEx(hFile HANDLE) (string, error) {

	var bufSize, err = GetFileSizeEx(hFile)

	if err != nil {
		panic(err)
	}

	lpszLongPath := make([]uint16, bufSize)

	ret, _, err := procReadFileEx.Call(
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpszLongPath[0])),
		uintptr(bufSize),
		uintptr(unsafe.Pointer(&bufSize)),
		0,
		0,
	)

	if ret == 0 {
		return "", err
	}

	return syscall.UTF16ToString(lpszLongPath), nil
}

//SetFileTime Sets the date and time that the specified file or directory was created, last accessed, or last modified.
func SetFileTime(hFile HANDLE, lpCreationTime, lpLastAccessTime, lpLastWriteTime FILETIME) error {
	ret, _, err := procSetFileTime.Call(
		hFile.ToUintPtr(),
		lpCreationTime.ToUintPtr(),
		lpLastAccessTime.ToUintPtr(),
		lpLastWriteTime.ToUintPtr(),
	)
	if ret == 0 {
		return err
	}

	return nil
}

//WriteFile Writes data to the specified file or input/output (I/O) device.
//
//This function is designed for both synchronous and asynchronous operation. For a similar function designed solely for asynchronous operation, see WriteFileEx.
func WriteFile(hFile HANDLE, data string) error {

	lpBuffer, err := syscall.UTF16FromString(data)
	if err != nil {
		return err
	}

	var buffer uint32

	ret, _, err := procWriteFile.Call(
		hFile.ToUintPtr(),
		uintptr(unsafe.Pointer(&lpBuffer)),
		uintptr(len(lpBuffer)),
		uintptr(unsafe.Pointer(&buffer)),
		0,
	)

	if ret == 0 {
		return err
	}

	return nil
}
