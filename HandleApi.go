package GoWinApi

import "syscall"

var (
	procCloseHandle          = kernel32.NewProc("CloseHandle")
	procCompareObjectHandles = kernel32.NewProc("CompareObjectHandles")
)

//CloseHandle Closes an open object handle.
func CloseHandle(handle syscall.Handle) (err error) {
	r, _, err := procCloseHandle.Call(uintptr(handle))
	if r == 0 {
		return
	}
	return nil
}

//CompareObjectHandles Compares two object handles to determine if they refer to the same underlying kernel object.
func CompareObjectHandles(firstObjectHandel, secondObjectHandel syscall.Handle) bool {
	r, _, _ := procCompareObjectHandles.Call(
		uintptr(firstObjectHandel),
		uintptr(secondObjectHandel),
	)

	return r != 0
}
