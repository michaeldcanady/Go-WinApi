package winbase

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"github.com/michaeldcanady/Go-WinApi/advapi32/sddl"
)

const (
	SidTypeUser           = 1
	SidTypeGroup          = 2
	SidTypeDomain         = 3
	SidTypeAlias          = 4
	SidTypeWellKnownGroup = 5
	SidTypeDeletedAccount = 6
	SidTypeInvalid        = 7
	SidTypeUnknown        = 8
	SidTypeComputer       = 9
	SidTypeLabel          = 10
	//SidTypeLogonSession
)

var (
	advapi32               = syscall.NewLazyDLL("advapi32.dll")
	procLookupAccountNameW = advapi32.NewProc("LookupAccountNameW")
)

func LookupAccountNameW(systemName, computerName, userName string) (string, string) {

	var (
		cbSid                   = 0
		cchReferencedDomainName = 0
		peUse                   uint16
		userNameInt, _          = syscall.UTF16FromString(userName)
	)

	_, _, _ = procLookupAccountNameW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(&userNameInt)),
		uintptr(0),
		uintptr(unsafe.Pointer(&cbSid)),
		uintptr(0),
		uintptr(unsafe.Pointer(&cchReferencedDomainName)),
		uintptr(unsafe.Pointer(&peUse)),
	)

	var Sid *syscall.SID
	var ReferencedDomainName int

	ret, _, err := procLookupAccountNameW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(&userNameInt)),
		uintptr(unsafe.Pointer(&Sid)),
		uintptr(unsafe.Pointer(&cbSid)),
		uintptr(unsafe.Pointer(&ReferencedDomainName)),
		uintptr(unsafe.Pointer(&cchReferencedDomainName)),
		uintptr(unsafe.Pointer(&peUse)),
	)

	if ret == 0 {
		fmt.Println(err)
	}

	SID, err := sddl.ConvertSidToStringSidW(Sid)

	SID = strings.Replace(SID, "-0", "", -1)

	if err != nil {
		return "", ""
	}

	fmt.Println("string", ReferencedDomainName)

	var PSID_NAME string

	switch peUse {
	case SidTypeUser:
		PSID_NAME = "User"
	case SidTypeGroup:
		PSID_NAME = "Group"
	case SidTypeDomain:
		PSID_NAME = "Domain"
	case SidTypeAlias:
		PSID_NAME = "Alias"
	case SidTypeWellKnownGroup:
		PSID_NAME = "WellKnownGroup"
	case SidTypeDeletedAccount:
		PSID_NAME = "DeletedAccount"
	case SidTypeInvalid:
		PSID_NAME = "Invalid"
	case SidTypeUnknown:
		PSID_NAME = "Unknown"
	case SidTypeComputer:
		PSID_NAME = "Computer"
	case SidTypeLabel:
		PSID_NAME = "Label"
	}

	return SID, PSID_NAME
}
