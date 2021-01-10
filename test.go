package main

import (
	"fmt"
	"strings"

	winapi "github.com/michaeldcanady/windows-api/windows-api"
	"github.com/michaeldcanady/windows-api/windows-api/kernel32"
)

func main() {
	//err := winapi.WNetAddConnection2("\\\\path\\to\\server", "username", "password")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("Connection successful")
	//}
	//address := "\\\\?\\Volume{8B45C404-2738-4BC6-8A55-A2E401D87A35}\\"
	//ret, err := kernel32.GetVolumeNameForVolumeMountPointW("C:\\")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ret)

	//fmt.Println(kernel32.GetVolumeInformationW("Z:\\"))
	//fmt.Println(kernel32.GetDriveTypeW())
	winapi.WNetOpenEnumW()
	//winapi.WNetEnumResourceW(handle)
	//drive()
	//kernel32.DeleteVolumeMountPointW("D:\\")
	//winapi.WNetCancelConnection("\\\\fs3.liberty.edu\\hdbackups", true)
}

func drive() {
	drives, err := kernel32.GetLogicalDrives()
	if err != nil {
		panic(err)
	}
	var remoteDrives []string
	for _, drive := range drives {
		drive = strings.ToUpper(drive)

		if !strings.HasSuffix(drive, ":\\") {
			drive = drive + ":\\"
		}

		ty, err := kernel32.GetDriveTypeW(drive)
		if err != nil {
			panic(err)
		}
		if ty == "remote" {
			remoteDrives = append(remoteDrives, drive)
		}
	}
	fmt.Println(remoteDrives)
}
