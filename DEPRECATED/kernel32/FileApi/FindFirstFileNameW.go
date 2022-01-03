package fileapi

import (
	"fmt"
	"unsafe"
)

//TODO Figure out how to get string length to work
func FindFirstFileNameW(fileName string) {

	var LinkName uintptr
	var stringLength uint32

	ret, _, err := procFindFirstFileNameW.Call(
		UintptrFromString(fileName),
		0,
		uintptr(unsafe.Pointer(&stringLength)),
		LinkName,
	)

	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println(LinkName)
}
