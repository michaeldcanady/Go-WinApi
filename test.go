package main

import (
	"fmt"
	"strings"

	fileapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/windows-api/kernel32/FileApi"
)

func main() {
	//fmt.Println(fileapi.GetVolumeInformationW("C:"))
	//handle, err := fileapi.CreateFileW(`C:\New folder\a.txt`,
	//	syscall.GENERIC_READ,
	//	syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
	//	0,
	//	syscall.CREATE_NEW,
	//	syscall.FILE_ATTRIBUTE_NORMAL,
	//	0)
	//if handle == syscall.InvalidHandle {
	//	fmt.Println(handle)
	//	fmt.Println(err)
	//}

	//err = handleapi.CloseHandle(handle)
	//if err != nil {
	//	log.Fatalf("second close: %v", err)
	//}
