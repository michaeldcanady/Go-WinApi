package fileapi

import "syscall"

type (
	LPVOID uintptr
	// DWORD dword typ
	DWORD   uint32
	LPBYTE  *byte
	PBYTE   *byte
	LPDWORD *uint32
	LPWSTR  *uint16
	LPCWSTR *uint16
)

type HANDLE syscall.Handle

type SecurityAttribute int64

func (H HANDLE) ToUintPtr() uintptr {
	return uintptr(H)
}
