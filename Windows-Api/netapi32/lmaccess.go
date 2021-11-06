package netapi32

import (
	"fmt"
	"time"
	"unsafe"

	fileapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/Windows-Api/kernel32/FileApi"
)

const (
	//User filters
	USER_FILTER_TEMP_DUPLICATE_ACCOUNT    = 0x0001
	USER_FILTER_NORMAL_ACCOUNT            = 0x0002
	USER_FILTER_INTERDOMAIN_TRUST_ACCOUNT = 0x0008
	USER_FILTER_WORKSTATION_TRUST_ACCOUNT = 0x0010
	USER_FILTER_SERVER_TRUST_ACCOUNT      = 0x0020

	USER_MAX_PREFERRED_LENGTH = 0xFFFFFFFF
)

var (
	ACCOUNT_CONTROL_FLAGS = map[int64]string{
		0x0001:    "UF_SCRIPT",
		0x0002:    "UF_ACCOUNTDISABLE",
		0x0008:    "UF_HOMEDIR_REQUIRED",
		0x0020:    "UF_PASSWD_NOTREQD",
		0x0040:    "UF_PASSWD_CANT_CHANGE",
		0x0010:    "UF_LOCKOUT",
		0x10000:   "UF_DONT_EXPIRE_PASSWD",
		0x0080:    "UF_ENCRYPTED_TEXT_PASSWORD_ALLOWED",
		0x100000:  "UF_NOT_DELEGATED",
		0x40000:   "UF_SMARTCARD_REQUIRED",
		0x200000:  "UF_USE_DES_KEY_ONLY",
		0x400000:  "UF_DONT_REQUIRE_PREAUTH",
		0x80000:   "UF_TRUSTED_FOR_DELEGATION",
		0x800000:  "UF_PASSWORD_EXPIRED",
		0x1000000: "UF_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION",
	}
)

// ListLocalUsers lists information about local user accounts.
func ListLocalUsers(serverName string, level int64, access uint32) ([]LocalUser1, error) {
	var (
		dataPointer  uintptr
		resumeHandle uintptr
		entriesRead  uint32
		entriesTotal uint32
		sizeTest     USER_INFO_1
		retVal       = make([]LocalUser1, 0)
	)

	ret, _, _ := usrNetUserEnum.Call(
		uintptr(0),                                 // servername
		uintptr(uint32(level)),                     // level, USER_INFO_2
		uintptr(uint32(access)),                    // filter, only "normal" accounts.
		uintptr(unsafe.Pointer(&dataPointer)),      // struct buffer for output data.
		uintptr(uint32(USER_MAX_PREFERRED_LENGTH)), // allow as much memory as required.
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
		var data = (*USER_INFO_1)(unsafe.Pointer(iter))

		ud := LocalUser1{
			Username:      UTF16toString(data.usri1_name),
			PasswordAge:   (time.Duration(data.usri1_password_age) * time.Second),
			HomeDirectory: UTF16toString(data.usri1_home_dir),
			Comment:       UTF16toString(data.usri1_comment),
			Flags:         fileapi.SeperateFlags(data.usri1_flags, ACCOUNT_CONTROL_FLAGS),
			ScriptPath:    UTF16toString(data.usri1_script_path),
		}

		if data.usri1_priv == USER_PRIV_GUEST {
			ud.Priviledge = "Guest"
		}
		if data.usri1_priv == USER_PRIV_USER {
			ud.Priviledge = "User"
		}
		if data.usri1_priv == USER_PRIV_ADMIN {
			ud.Priviledge = "Admin"
		}

		retVal = append(retVal, ud)
		iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
	}

	usrNetApiBufferFree.Call(dataPointer)
	return retVal, nil
}

//if (data.Usri2_flags & USER_UF_ACCOUNTDISABLE) != USER_UF_ACCOUNTDISABLE {
//	ud.IsEnabled = true
//}
//if (data.Usri2_flags & USER_UF_LOCKOUT) == USER_UF_LOCKOUT {
//	ud.IsLocked = true
//}
//if (data.Usri2_flags & USER_UF_PASSWD_CANT_CHANGE) == USER_UF_PASSWD_CANT_CHANGE {
//	ud.NoChangePassword = true
//}
//if (data.Usri2_flags & USER_UF_DONT_EXPIRE_PASSWD) == USER_UF_DONT_EXPIRE_PASSWD {
//	ud.PasswordNeverExpires = true
//}
//if data.Usri2_priv == USER_PRIV_ADMIN {
//	ud.IsAdmin = true
//}

//retVal = append(retVal, ud)

//iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
//}
