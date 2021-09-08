package handleapi

import "syscall"

var (
	procCloseHandle = kernel32.NewProc("CloseHandle")
)

func CloseHandle(handle syscall.Handle) error {
	r, _, err := procCloseHandle.Call(uintptr(handle))
	if r == 0 {
		return err
	}
	return nil
}
