package fileapi

import (
	"syscall"
	"unsafe"
)

//ReadFileEx Reads data from the specified file or input/output (I/O) device. It reports its completion status asynchronously, calling the specified completion
//routine when reading is completed or canceled and the calling thread is in an alertable wait state.
func ReadFileEx(hFile HANDLE) (string, error) {

	var bufSize, err = GetFileSize(hFile)

	if err != nil {
		panic(err)
	}

	lpszLongPath := make([]uint16, bufSize)

	//var ran uint32

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
