package main

import (
	"fmt"

	winapi "github.com/michaeldcanady/Windows-Api/Windows-Api"
)

func main() {
	err := winapi.WNetAddConnection2("\\\\path\\to\\server", "username", "password")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection successful")
	}
}
