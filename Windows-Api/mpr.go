package winapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

type (
	DWORD  uint32
	LPTSTR uintptr
)

type cError struct {
	hex   int
	error string
}

func (e *cError) Error() string {
	return fmt.Sprintf("Your entries raised a %s", e.error)
}

const (
	//Errors
	WN_SUCCESS  = NO_ERROR
	WN_NO_ERROR = NO_ERROR

	//Resource types
	RESOURCETYPE_ANY   = 0x00000000
	RESOURCETYPE_DISK  = 0x00000001
	RESOURCETYPE_PRINT = 0x00000002
	// Connection Flags
	CONNECT_UPDATE_PROFILE = 0x00000001
	CONNECT_UPDATE_RECENT  = 0x00000002
	CONNECT_TEMPORARY      = 0x00000004
	CONNECT_INTERACTIVE    = 0x00000008
	CONNECT_PROMPT         = 0x00000010
	CONNECT_REDIRECT       = 0x00000080
	CONNECT_CMD_SAVECRED   = 0x00001000
	CONNECT_CRED_RESET     = 0x00002000

	RESOURCE_CONNECTED               = 0x00000001
	RESOURCE_GLOBALNET               = 0x00000002
	RESOURCE_RECENT                  = 0x00000004
	RESOURCE_CONTEXT                 = 0x00000005
	RESOURCETYPE_RESERVED            = 0x00000008
	RESOURCETYPE_UNKNOWN             = 0xFFFFFFFF
	RESOURCEUSAGE_CONNECTABLE        = 0x00000001
	RESOURCEUSAGE_CONTAINER          = 0x00000002
	RESOURCEUSAGE_NOLOCALDEVICE      = 0x00000004
	RESOURCEUSAGE_SIBLING            = 0x00000008
	RESOURCEUSAGE_ATTACHED           = 0x00000010
	RESOURCEUSAGE_RESERVED           = 0x80000000
	RESOURCEDISPLAYTYPE_GENERIC      = 0x00000000
	RESOURCEDISPLAYTYPE_DOMAIN       = 0x00000001
	RESOURCEDISPLAYTYPE_SERVER       = 0x00000002
	RESOURCEDISPLAYTYPE_SHARE        = 0x00000003
	RESOURCEDISPLAYTYPE_FILE         = 0x00000004
	RESOURCEDISPLAYTYPE_GROUP        = 0x00000005
	RESOURCEDISPLAYTYPE_NETWORK      = 0x00000006
	RESOURCEDISPLAYTYPE_ROOT         = 0x00000007
	RESOURCEDISPLAYTYPE_SHAREADMIN   = 0x00000008
	RESOURCEDISPLAYTYPE_DIRECTORY    = 0x00000009
	RESOURCEDISPLAYTYPE_TREE         = 0x0000000A
	RESOURCEDISPLAYTYPE_NDSCONTAINER = 0x0000000B
)

var (
	mpr, _                      = syscall.LoadDLL("mpr.dll")
	procWNetAddConnection2, _   = mpr.FindProc("WNetAddConnection2W")
	procWNetAddConnection, _    = mpr.FindProc("WNetAddConnectionW")
	procWNetCancelConnection, _ = mpr.FindProc("WNetCancelConnectionW")
	procWNetOpenEnum, _         = mpr.FindProc("WNetOpenEnumW")
)

type NETRESOURCE struct {
	dwScope       DWORD
	dwType        DWORD
	dwDisplayType DWORD
	dwUsage       DWORD
	lpLocalName   LPTSTR
	lpRemoteName  LPTSTR
	lpComment     LPTSTR
	lpProvider    LPTSTR
}

func CreateNetResource(dw int, remote, localName string) NETRESOURCE {
	path := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(remote)))
	local := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(localName)))
	n := NETRESOURCE{dwType: DWORD(dw), lpRemoteName: LPTSTR(path), lpLocalName: LPTSTR(local)}
	return n
}

func WNetOpenEnumW() {
	var resource NETRESOURCE
	var handle uintptr

	ret, _, err := procWNetOpenEnum.Call(
		RESOURCE_GLOBALNET,
		RESOURCETYPE_ANY,
		0,
		uintptr(unsafe.Pointer(&resource)),
		handle,
	)

	fmt.Println(ret, err, handle)
}

// More detailed information at: https://docs.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetaddconnectionw
func WNetAddConnection(remote, password, localName string) error {
	lpRemoteName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(remote)))
	lpPassword := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(password)))
	lpLocalName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(localName)))

	ret, _, _ := procWNetAddConnection.Call(lpRemoteName, lpPassword, lpLocalName)

	switch ret {
	case NO_ERROR, ERROR_SESSION_CREDENTIAL_CONFLICT: //succesfully authenticate or already did
		return nil
	case 0x43:
		return fmt.Errorf("Please verify server address.")
	case 0x35:
		return fmt.Errorf("Please set a user.")
	case 0x56:
		return fmt.Errorf("Please set a password")
	default:
		return fmt.Errorf("Error has appeared, return value %x.\n Please create an issue.", ret)
	}
}

// More detailed information at: https://docs.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetaddconnection2w
func WNetAddConnection2(path, user, pass, volume string) error {
	resource := CreateNetResource(RESOURCETYPE_DISK, path, volume)
	password := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(pass)))
	username := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(user)))

	ret, _, _ := procWNetAddConnection2.Call(uintptr(unsafe.Pointer(&resource)), password, username, CONNECT_TEMPORARY)

	switch ret {
	case NO_ERROR, ERROR_SESSION_CREDENTIAL_CONFLICT: //succesfully authenticate or already did
		return nil
	case 0x43:
		return fmt.Errorf("Please verify server address.")
	case 0x35:
		return fmt.Errorf("Please set a user.")
	case 0x56:
		return fmt.Errorf("Please set a password")
	default:
		return fmt.Errorf("Error has appeared, return value %x.\n Please create an issue.", ret)
	}
}

// More detailed information at: https://docs.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetcancelconnectionw
//func WNetCancelConnection(name string, force bool) error {

//	lpName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))
//	force1 := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(force)))

//	ret, _, _ := WNetCancelConnection1.Call(lpName, force1)

//	switch ret {
//	case NO_ERROR, ERROR_SESSION_CREDENTIAL_CONFLICT: //succesfully authenticate or already did
//		return nil
//	case 0x43:
//		return fmt.Errorf("Please verify server address.")
//	case 0x35:
//		return fmt.Errorf("Please set a user.")
//	case 0x56:
//		return fmt.Errorf("Please set a password")
//	default:
//		return fmt.Errorf("Error has appeared, return value %x.\n Please create an issue.", ret)
//	}
//}
