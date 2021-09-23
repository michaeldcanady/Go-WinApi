package main

import (

	//"strings"
	//"syscall"

	"fmt"
	"syscall"

	fileapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/FileApi"
)

func main() {

	handle, err := fileapi.CreateFileW(`C:\Users\micha\OneDrive\Documents\Villum's Story.docx`, syscall.GENERIC_READ,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		0,
		syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS,
		0)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileapi.ReadFileEx(handle))

	//handle, data, _ := fileapi.FindFirstStreamW(`C:\Users\micha\OneDrive\Documents\backup.exe`)
	//fmt.Println(data)
	//for {
	//	_, data, err := fileapi.FindNextStreamW(handle)
	//	if fmt.Sprintf("%s", err) == "Reached the end of the file." {
	//		err = fileapi.FindClose(handle)
	//		if err != nil {
	//			panic(err)
	//		}
	//		break
	//	}
	//	fmt.Println(data)
	//}

	//fmt.Println(fileapi.AreFileApisANSI())

	//handle, filedata, err := fileapi.FindFirstFileW(`C:\Users\micha\OneDrive\Documents\*.doc`)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(filedata)

	//for {
	//	data, err := fileapi.FindNextFileW(handle)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(data)
	//}

	//fmt.Println(fileapi.GetVolumeInformationW("C:"))
	//if handle == syscall.InvalidHandle {
	//	fmt.Println(handle)
	//	fmt.Println(err)
	//}

	//err = handleapi.CloseHandle(handle)
	//if err != nil {
	//	log.Fatalf("second close: %v", err)
	//}

	//fmt.Println(fileapi.FindFirstFileExW(`C:\Users\micha\OneDrive\Documents\*`,
	//	fileapi.FindExInfoStandard,
	//	fileapi.FindExSearchNameMatch,
	//	0))
}
