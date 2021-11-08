package main

import (
	"fmt"

	"github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/advapi32/winbase"
	"golang.org/x/sys/windows/registry"
)

//"strings"
//"syscall"

func main() {

	//handle, err := fileapi.CreateFileW(`Filepath\file.txt`, syscall.GENERIC_READ,
	//	syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
	//	0,
	//	syscall.OPEN_EXISTING,
	//	syscall.FILE_FLAG_BACKUP_SEMANTICS,
	//	0)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(fileapi.WriteFile(handle, "String"))

	//fmt.Println(netapi32.ListLocalUsers("test", 1, netapi32.USER_FILTER_NORMAL_ACCOUNT))
	//fmt.Println(sysinfoapi.EnumSystemFirmwareTables(sysinfoapi.RSMB))
	SID, _ := winbase.LookupAccountNameW("", "DESKTOP-FJ8V04D", "michael")

	fmt.Println(SID)

	fmt.Println(findHomeDirInRegistry(SID))
	//Ssysinfoapi.GetSystemFirmwareTable(sysinfoapi.ACPI, sysinfoapi.MSDM)

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

func findHomeDirInRegistry(uid string) (dir string, e error) {
	k, e := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\ProfileList\`+uid, registry.QUERY_VALUE)
	if e != nil {
		return "", e
	}
	defer k.Close()
	dir, _, e = k.GetStringValue("ProfileImagePath")
	if e != nil {
		return "", e
	}
	return dir, nil
}
