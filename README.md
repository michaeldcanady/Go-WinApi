# Windows-Api
## About Windows-API
Other more technical versions exist. The aim of this one is to make a way that is simple to use and you don't need much knowledge of C++ to get started.
## Gettings Started
1. Make sure you have a [GCC installed](#recommended-way-to-install-gcc)

## Contributing
    

## Recommended way to install GCC

## API:
FileAPI
### functions:

AreFileApisANSI

CompareFileTime

~~CreateDirectoryW~~

CreateFile2

~~CreateFileW~~

DefineDosDeviceW

~~DeleteFileW~~

DeleteVolumeMountPointW

FileTimeToLocalFileTime

FindClose

FindCloseChangeNotification

FindFirstChangeNotificationA

FindFirstChangeNotificationW

FindFirstFileA

FindFirstFileExA

FindFirstFileExW

FindFirstFileNameW

FindFirstFileW

FindFirstStreamW

FindFirstVolumeW

FindNextChangeNotification

FindNextFileA

Continues a file search from a previous call to the FindFirstFile, FindFirstFileEx, or FindFirstFileTransacted functions.
FindNextFileNameW

Continues enumerating the hard links to a file using the handle returned by a successful call to the FindFirstFileNameW function.
FindNextFileW

Continues a file search from a previous call to the FindFirstFile, FindFirstFileEx, or FindFirstFileTransacted functions.
FindNextStreamW

Continues a stream search started by a previous call to the FindFirstStreamW function.
FindNextVolumeW

Continues a volume search started by a call to the FindFirstVolume function.
FindVolumeClose

Closes the specified volume search handle.
FlushFileBuffers

Flushes the buffers of a specified file and causes all buffered data to be written to a file.
GetCompressedFileSizeA

Retrieves the actual number of bytes of disk storage used to store a specified file.
GetCompressedFileSizeW

Retrieves the actual number of bytes of disk storage used to store a specified file.
GetDiskFreeSpaceA

Retrieves information about the specified disk, including the amount of free space on the disk.
GetDiskFreeSpaceExA

Retrieves information about the amount of space that is available on a disk volume, which is the total amount of space, the total amount of free space, and the total amount of free space available to the user that is associated with the calling thread.
GetDiskFreeSpaceExW

Retrieves information about the amount of space that is available on a disk volume, which is the total amount of space, the total amount of free space, and the total amount of free space available to the user that is associated with the calling thread.
GetDiskFreeSpaceW

Retrieves information about the specified disk, including the amount of free space on the disk.
GetDriveTypeA

Determines whether a disk drive is a removable, fixed, CD-ROM, RAM disk, or network drive.
GetDriveTypeW

Determines whether a disk drive is a removable, fixed, CD-ROM, RAM disk, or network drive.
GetFileAttributesA

Retrieves file system attributes for a specified file or directory.
GetFileAttributesExA

Retrieves attributes for a specified file or directory.
GetFileAttributesExW

Retrieves attributes for a specified file or directory.
GetFileAttributesW

Retrieves file system attributes for a specified file or directory.
GetFileInformationByHandle

Retrieves file information for the specified file.
GetFileSize

Retrieves the size of the specified file, in bytes.
GetFileSizeEx

Retrieves the size of the specified file.
GetFileTime

Retrieves the date and time that a file or directory was created, last accessed, and last modified.
GetFileType

Retrieves the file type of the specified file.
GetFinalPathNameByHandleA

Retrieves the final path for the specified file.
GetFinalPathNameByHandleW

Retrieves the final path for the specified file.
GetFullPathNameA

Retrieves the full path and file name of the specified file.
GetFullPathNameW

Retrieves the full path and file name of the specified file.
GetLogicalDrives

Retrieves a bitmask representing the currently available disk drives.
GetLogicalDriveStringsW

Fills a buffer with strings that specify valid drives in the system.
GetLongPathNameA

Converts the specified path to its long form.
GetLongPathNameW

Converts the specified path to its long form.
GetShortPathNameW

Retrieves the short path form of the specified path.
GetTempFileNameA

Creates a name for a temporary file. If a unique file name is generated, an empty file is created and the handle to it is released; otherwise, only a file name is generated.
GetTempFileNameW

Creates a name for a temporary file. If a unique file name is generated, an empty file is created and the handle to it is released; otherwise, only a file name is generated.
GetTempPath2A

Retrieves the path of the directory designated for temporary files, based on the privileges of the calling process.
GetTempPath2W

Retrieves the path of the directory designated for temporary files, based on the privileges of the calling process.
GetTempPathA

Retrieves the path of the directory designated for temporary files.
GetTempPathW

Retrieves the path of the directory designated for temporary files.
GetVolumeInformationA

Retrieves information about the file system and volume associated with the specified root directory.
GetVolumeInformationByHandleW

Retrieves information about the file system and volume associated with the specified file.
GetVolumeInformationW

Retrieves information about the file system and volume associated with the specified root directory.
GetVolumeNameForVolumeMountPointW

Retrieves a volume GUID path for the volume that is associated with the specified volume mount point ( drive letter, volume GUID path, or mounted folder).
GetVolumePathNamesForVolumeNameW

Retrieves a list of drive letters and mounted folder paths for the specified volume.
GetVolumePathNameW

Retrieves the volume mount point where the specified path is mounted.
LocalFileTimeToFileTime

Converts a local file time to a file time based on the Coordinated Universal Time (UTC).
LockFile

Locks the specified file for exclusive access by the calling process.
LockFileEx

Locks the specified file for exclusive access by the calling process. This function can operate either synchronously or asynchronously and can request either an exclusive or a shared lock.
QueryDosDeviceW

Retrieves information about MS-DOS device names.
ReadFile

Reads data from the specified file or input/output (I/O) device. Reads occur at the position specified by the file pointer if supported by the device.
ReadFileEx

Reads data from the specified file or input/output (I/O) device. It reports its completion status asynchronously, calling the specified completion routine when reading is completed or canceled and the calling thread is in an alertable wait state.
ReadFileScatter

Reads data from a file and stores it in an array of buffers.
RemoveDirectoryA

Deletes an existing empty directory.
RemoveDirectoryW

Deletes an existing empty directory.
SetEndOfFile

Sets the physical file size for the specified file to the current position of the file pointer.
SetFileApisToANSI

Causes the file I/O functions to use the ANSI character set code page for the current process.
SetFileApisToOEM

Causes the file I/O functions for the process to use the OEM character set code page.
SetFileAttributesA

Sets the attributes for a file or directory.
SetFileAttributesW

Sets the attributes for a file or directory.
SetFileInformationByHandle

Sets the file information for the specified file.
SetFileIoOverlappedRange

Associates a virtual address range with the specified file handle.
SetFilePointer

Moves the file pointer of the specified file.
SetFilePointerEx

Moves the file pointer of the specified file.
SetFileTime

Sets the date and time that the specified file or directory was created, last accessed, or last modified.
SetFileValidData

Sets the valid data length of the specified file. This function is useful in very limited scenarios. For more information, see the Remarks section.
UnlockFile

Unlocks a region in an open file.
UnlockFileEx

Unlocks a region in the specified file. This function can operate either synchronously or asynchronously.
WriteFile

Writes data to the specified file or input/output (I/O) device.
WriteFileEx

Writes data to the specified file or input/output (I/O) device. It reports its completion status asynchronously, calling the specified completion routine when writing is completed or canceled and the calling thread is in an alertable wait state.
WriteFileGather

Retrieves data from an array of buffers and writes the data to a file.