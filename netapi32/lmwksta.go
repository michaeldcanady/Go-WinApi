package netapi32

import (
	"syscall"
	"unsafe"
)

var (
	modNetapi32         = syscall.NewLazyDLL("netapi32.dll")
	usrNetUserEnum      = modNetapi32.NewProc("NetUserEnum")
	usrNetApiBufferFree = modNetapi32.NewProc("NetApiBufferFree")
)

// UTF16toString converts a pointer to a UTF16 string into a Go string.
func UTF16toString(p *uint16) string {
	return syscall.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}
