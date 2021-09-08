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

	//fmt.Println(fileapi.CreateDirectoryW(`C:\New folder\testfolder`, 0))
	fmt.Println(fileapi.DeleteFileW(`C:\New folder\a.txt`))
}

func drive() {
	drives, err := fileapi.GetLogicalDrives()
	if err != nil {
		panic(err)
	}
	for _, drive := range drives {
		drive = strings.ToUpper(drive)

		if !strings.HasSuffix(drive, ":\\") {
			drive = drive + ":\\"
		}

		ty, err := fileapi.GetDriveTypeW(drive)
		if err != nil {
			panic(err)
		}
		fmt.Println(ty)
	}
}
