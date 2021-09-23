package fileapi

import (
	"errors"
	"fmt"
	"syscall"
)

var (
	createFileW = kernel32.NewProc("CreateFileW")

	shareMode = map[string]int{
		"DO_NOT_FILE_SHARE": 0x00000000,
		"FILE_SHARE_DELETE": 0x00000004,
		"FILE_SHARE_READ":   0x00000001,
		"FILE_SHARE_WRITE":  0x00000002,
	}
	creationDisposition = map[string]int{
		"CREATE_NEW":        1,
		"CREATE_ALWAYS":     2,
		"OPEN_EXISTING":     3,
		"OPEN_ALWAYS":       4,
		"TRUNCATE_EXISTING": 5,
	}
	FlagsAndAttributes = map[string]int{
		"FILE_ATTRIBUTE_ARCHIVE":   32,
		"FILE_ATTRIBUTE_ENCRYPTED": 16384,
		"FILE_ATTRIBUTE_HIDDEN":    2,
		"FILE_ATTRIBUTE_NORMAL":    128, //Only Used Alone
		"FILE_ATTRIBUTE_OFFLINE":   4096,
		"FILE_ATTRIBUTE_READONLY":  1,
		"FILE_ATTRIBUTE_SYSTEM":    4,
		"FILE_ATTRIBUTE_TEMPORARY": 256,
	}
)

//CreateFileW Creates or opens a file or I/O device. The most commonly used I/O devices are as follows: file, file stream, directory, physical disk, volume, console buffer, tape drive, communications resource, mailslot, and pipe. The function returns a handle that can be used to access the file or device for various types of I/O depending on the file or device and the flags and attributes specified.
//To perform this operation as a transacted operation, which results in a handle that can be used for transacted I/O, use the CreateFileTransacted function.
func CreateFileW(FileName string, dwDesiredAccess DWORD, dwShareMode DWORD, lpSecurityAttributes interface{}, dwCreationDisposition DWORD, dwFlagsAndAttributes DWORD, dhTemplateFile syscall.Handle) (syscall.Handle, error) {

	if lpSecurityAttributes.(int) != 0 && fmt.Sprintf("%T", lpSecurityAttributes) != "*_SECURITY_ATTRIBUTES" {
		return syscall.Handle(0), errors.New("Incompatible lpSecurityAttributes used please use 0 or specify security attributes")
	}

	ret, _, err := createFileW.Call(
		UintptrFromString(&FileName),        // [in] LPCTSTR
		uintptr(dwDesiredAccess),            // [in] DWORD
		uintptr(dwShareMode),                // [in] DWORD
		uintptr(lpSecurityAttributes.(int)), // [in] LPSECURITY_ATT...
		uintptr(dwCreationDisposition),      // [in] DWORD
		uintptr(dwFlagsAndAttributes),       // [in] DWORD
		uintptr(dhTemplateFile),             // [in] HANDLE
	)

	if ret == 0 {
		return syscall.Handle(0), err
	}

	return syscall.Handle(ret), nil
}
