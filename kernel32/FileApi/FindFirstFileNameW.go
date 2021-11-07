package fileapi

import(
	"fmt"
	"unsafe"
)

var findFirstFileNameWProc = kernel32.NewProc("FindFirstFileNameW")


//TODO Figure out how to get string length to work
func FindFirstFileNameW(lpFileName string){

	var LinkName string
	var stringlength *uint32

	ret,_,err := findFirstFileNameWProc.Call(
		UintptrFromString(&lpFileName),
		0,
		uintptr(unsafe.Pointer(stringlength)),
		UintptrFromString(&LinkName),
	)

	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println(LinkName)
}