package fileapi

import (
	"syscall"
	"unsafe"
)

func WriteFile(hFile HANDLE, data string) (bitsWritten int64, err error) {

	lpBuffer, err := syscall.UTF16FromString(data)
	if err != nil {
		return bitsWritten, err
	}

	ret, _, err := writeFileProc.Call(
		hFile.toUTF16Ptr(),                    //IN hFile
		uintptr(unsafe.Pointer(&lpBuffer)),    // IN lpBuffer
		uintptr(len(lpBuffer)),                // IN nNumberOfBytesToWrite
		uintptr(unsafe.Pointer(&bitsWritten)), // OUT, OPTIONAL lpNumberOfBytesWritten
		0,                                     // IN, OUT, OPTIONAL lpOverlapped
	)

	if ret == 0 {
		return bitsWritten, err
	}

	return bitsWritten, nil
}
