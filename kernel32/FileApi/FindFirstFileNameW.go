package fileapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

//TODO Figure out how to get string length to work
func FindFirstFileNameW(fileName string) {

	var LinkName string
	var stringlength *uint32

	lpFileName, err := syscall.UTF16PtrFromString(fileName)

	if err != nil {
		fmt.Println(err)
	}

	lpLinkName, err := syscall.UTF16PtrFromString(LinkName)

	if err != nil {
		fmt.Println(err)
	}

	ret, _, err := findFirstFileNameWProc.Call(
		lpFileName,
		0,
		uintptr(unsafe.Pointer(stringlength)),
		lpLinkName,
	)

	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println(LinkName)
}
