package AreFileApisANSIExample

import (
	"fmt"

	gowinapi "github.com/michaeldcanady/go-winapi/gowinapi"
)

func main() {
	UsesOemCodePages := gowinapi.AreFileApisANSI()
	fmt.Println(UsesOemCodePages)
}
