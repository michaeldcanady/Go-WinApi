package netapi32

import (
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

type SessionInfo struct {
	Cname    string
	Username string
	Time     time.Duration
	IdleTime time.Duration
}

type sessionInfo10 struct {
	Cname    *uint16
	Username *uint16
	Time     uint32
	IdleTime uint32
}

var (
	modNetApi32 = syscall.NewLazyDLL("netapi32.dll")

	procNetApiBufferReallocate = modNetApi32.NewProc("NetApiBufferReallocate")
	procNetApiBufferAllocate   = modNetApi32.NewProc("NetApiBufferAllocate")
	procNetApiBufferFree       = modNetApi32.NewProc("NetApiBufferFree")
	procNetApiBufferSize       = modNetApi32.NewProc("NetApiBufferSize")
	procNetSessionEnum         = modNetApi32.NewProc("NetSessionEnum")
)

func NetApiBufferFree(p unsafe.Pointer) error {
	rc, _, _ := syscall.Syscall(procNetApiBufferFree.Addr(), 1, uintptr(p), 0, 0)
	if rc != 0 {
		return syscall.Errno(rc)
	}
	return nil
}

//NOT A WINDOWS-API CALL
func mustFreeNetApiBuffer(p unsafe.Pointer) {
	err := NetApiBufferFree(p)
	if err != nil {
		panic(err)
	}
}

func NetSessionEnum() (out []SessionInfo, err error) {
	var (
		pinfo        *sessionInfo10
		prefmaxlen   int32 = -1
		entriesread  uint32
		totalentries uint32
	)

	rc, _, _ := syscall.Syscall12(procNetSessionEnum.Addr(), 9,
		0, 0, 0, 10,
		uintptr(unsafe.Pointer(&pinfo)),
		uintptr(prefmaxlen),
		uintptr(unsafe.Pointer(&entriesread)),
		uintptr(unsafe.Pointer(&totalentries)),
		0,
		0, 0, 0)

	defer mustFreeNetApiBuffer(unsafe.Pointer(pinfo))

	if rc != 0 {
		err = syscall.Errno(rc)
		return
	}

	if entriesread == 0 {
		return
	}

	out = make([]SessionInfo, entriesread)
	for i := uint32(0); i < entriesread; i++ {
		out[i] = SessionInfo{
			Cname:    UTF16PtrToString(pinfo.Cname),
			Username: UTF16PtrToString(pinfo.Username),
			Time:     time.Duration(pinfo.Time) * time.Second,
			IdleTime: time.Duration(pinfo.IdleTime) * time.Second,
		}
		p := uintptr(unsafe.Pointer(pinfo))
		p += unsafe.Sizeof(*pinfo)
		pinfo = (*sessionInfo10)(unsafe.Pointer(p))
	}

	return
}

func NetApiBufferAllocate() {

}
