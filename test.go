package main

import (
	"fmt"
	//"strings"
	"syscall"

	fileapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/windows-api/kernel32/FileApi"
)

func main() {

	fmt.Println(fileapi.AreFileApisANSI())

	//handle, filedata, err := fileapi.FindFirstFileW(`C:\Users\micha\OneDrive\Documents\*.doc`)
	//if err != nil {
//		panic(err)
//	}
//	fmt.Println(filedata)

//	for{
//		data, err := fileapi.FindNextFileW(handle)
//		if err != nil {
//			break
//		}
//		fmt.Println(data)
//	}

	//fmt.Println(fileapi.GetVolumeInformationW("C:"))
	//if handle == syscall.InvalidHandle {
	//	fmt.Println(handle)
	//	fmt.Println(err)
	//}

	//err = handleapi.CloseHandle(handle)
	//if err != nil {
	//	log.Fatalf("second close: %v", err)
	//}
}