package winapi

import(
  "syscall"
	"unsafe"
  "golang.org/x/sys/windows"
  "fmt"
)

const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

var _ unsafe.Pointer

var (
	// Points to the desired dll or 'WINDOWS API'
  user32DLL,_           = syscall.LoadDLL("user32.dll")
	UserenvDLL          = windows.NewLazyDLL("Userenv.dll")
	Netapi32DLL         = windows.NewLazyDLL("Netapi32.dll")

  procSystemParamInfo,_  = user32DLL.FindProc("SystemParametersInfoW")
	procDeleteProfile       = UserenvDLL.NewProc("DeleteProfileW")
	NetUserEnum         = Netapi32DLL.NewProc("NetUserEnum")



	python = `C:\Users\dmcanady\Downloads\python.png`
	golang = `C:\Users\dmcanady\Downloads\golang.png`
)

// Removes user and User directory!!!
func DeleteProfile(SID string)error{
	UserSid := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(SID)))
	ret, _, _ := procDeleteProfile.Call(UserSid,0,0)
	switch ret{
  default:
    return fmt.Errorf("Error has appeared, return value %x.\n Please create an issue.", ret)
  }
}

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}
