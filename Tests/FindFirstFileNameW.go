package main

import "github.com/michaeldcanady/Go-WinApi/kernel32/FileApi"

func main() {
	fileapi.FindFirstFileNameW("C:\\Program Files (x86)\\Dell\\CCTK\\cctkerrorcodes.txt")
}
