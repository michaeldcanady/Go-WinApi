package GoWinApi

import "syscall"

var (
	procCloseHandle = kernel32.NewProc("CloseHandle")
)

//CloseHandle Closes an open object handle.
func CloseHandle(handle syscall.Handle) (err error) {
	r, _, err := procCloseHandle.Call(uintptr(handle))
	if r == 0 {
		return
	}
	return nil
}
