package fileapi

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
