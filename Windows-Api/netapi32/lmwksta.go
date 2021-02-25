package netapi32

import(
  "fmt"
	"syscall"
	"time"
	"unsafe"

	so "github.com/iamacarpet/go-win64api/shared"
)

var(
  modNetapi32                = syscall.NewLazyDLL("netapi32.dll")
  usrNetUserEnum             = modNetapi32.NewProc("NetUserEnum")
  usrNetApiBufferFree        = modNetapi32.NewProc("NetApiBufferFree")
)

// UTF16toString converts a pointer to a UTF16 string into a Go string.
func UTF16toString(p *uint16) string {
	return syscall.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}

// ListLocalUsers lists information about local user accounts.
func ListLocalUsers(server string) ([]so.LocalUser, error) {
	var (
		dataPointer  uintptr
		resumeHandle uintptr
		entriesRead  uint32
		entriesTotal uint32
		sizeTest     USER_INFO_2
		retVal       = make([]so.LocalUser, 0)
	)
  s16, err := syscall.UTF16PtrFromString(server)
  if err != nil {
    panic(err)
  }

	ret, _, _ := usrNetUserEnum.Call(
		uintptr(unsafe.Pointer(s16)),         // servername
		uintptr(uint32(2)), // level, USER_INFO_2
		uintptr(uint32(USER_FILTER_NORMAL_ACCOUNT)), // filter, only "normal" accounts.
		uintptr(unsafe.Pointer(&dataPointer)),       // struct buffer for output data.
		uintptr(uint32(USER_MAX_PREFERRED_LENGTH)),  // allow as much memory as required.
		uintptr(unsafe.Pointer(&entriesRead)),
		uintptr(unsafe.Pointer(&entriesTotal)),
		uintptr(unsafe.Pointer(&resumeHandle)),
	)
	if ret != NET_API_STATUS_NERR_Success {
		return nil, fmt.Errorf("error fetching user entry")
	} else if dataPointer == uintptr(0) {
		return nil, fmt.Errorf("null pointer while fetching entry")
	}

	var iter = dataPointer
	for i := uint32(0); i < entriesRead; i++ {
		var data = (*USER_INFO_2)(unsafe.Pointer(iter))

		ud := so.LocalUser{
			Username:         UTF16toString(data.Usri2_name),
			FullName:         UTF16toString(data.Usri2_full_name),
			PasswordAge:      (time.Duration(data.Usri2_password_age) * time.Second),
			LastLogon:        time.Unix(int64(data.Usri2_last_logon), 0),
			BadPasswordCount: data.Usri2_bad_pw_count,
			NumberOfLogons:   data.Usri2_num_logons,
		}

		if (data.Usri2_flags & USER_UF_ACCOUNTDISABLE) != USER_UF_ACCOUNTDISABLE {
			ud.IsEnabled = true
		}
		if (data.Usri2_flags & USER_UF_LOCKOUT) == USER_UF_LOCKOUT {
			ud.IsLocked = true
		}
		if (data.Usri2_flags & USER_UF_PASSWD_CANT_CHANGE) == USER_UF_PASSWD_CANT_CHANGE {
			ud.NoChangePassword = true
		}
		if (data.Usri2_flags & USER_UF_DONT_EXPIRE_PASSWD) == USER_UF_DONT_EXPIRE_PASSWD {
			ud.PasswordNeverExpires = true
		}
		if data.Usri2_priv == USER_PRIV_ADMIN {
			ud.IsAdmin = true
		}

		retVal = append(retVal, ud)

		iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
	}
	usrNetApiBufferFree.Call(dataPointer)
	return retVal, nil
}
