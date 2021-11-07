package fileapi

import (
	"syscall"
	"unsafe"
)

var getFinalPathNameByHandleWProc = kernel32.NewProc("GetFinalPathNameByHandleW")

func GetFinalPathNameByHandleW(hFile syscall.Handle) (string, string, error) {

	var bufSize uint32 = syscall.MAX_PATH // 260
	buf := make([]uint16, bufSize)
	var rawFlags uint32

	ret, _, err := getFinalPathNameByHandleWProc.Call(
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
