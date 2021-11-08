package fileapi

import (
	"syscall"
	"unsafe"
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

	return syscall.UTF16ToString(buf), getFlag(rawFlags), nil
}

func getFlag(rawFlags uint32) (flag string) {
	switch rawFlags {
	case 0x0:
		flag = "FILE_NAME_NORMALIZED"
	case 0x8:
		flag = "FILE_NAME_OPENED"
	}
	return
}
