package main

import (
	"fmt"

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
	ret, err := kernel32.GetVolumeNameForVolumeMountPointW("C:")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	//kernel32.GetVolumeInformationW("\\\\Share\\Share" or "C:\\")

	fmt.Println(kernel32.GetLogicalDrives())
}
