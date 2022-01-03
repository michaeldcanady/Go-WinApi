package fileapi

import (
	"syscall"
	"unsafe"
)

var (
	FileNameFlags = map[int64]string{
		0x0: "FILE_NAME_NORMALIZED",
		0x8: "FILE_NAME_OPENED",
	}
)

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
