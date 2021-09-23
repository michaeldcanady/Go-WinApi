package fileapi

import (
	"syscall"
	"unsafe"
)

var getFinalPathNameByHandleWProc = kernel32.NewProc("GetFinalPathNameByHandleW")

func GetFinalPathNameByHandleW(hFile syscall.Handle) (Name string, Flag string, err error) {

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
		Name = ""

	} else {
		Name = uint16ToString(buf)
		err = nil
	}

	switch rawFlags {
	case 0x0:
		Flag = "FILE_NAME_NORMALIZED"
	case 0x8:
		Flag = "FILE_NAME_OPENED"
	}

	return
}
