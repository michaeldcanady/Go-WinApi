package winapi

import (
	"unicode/utf16"
	"unsafe"
)

func Utf16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}

	s := make([]uint16, 0, 50)
	for {
		if *p == 0 {
			return string(utf16.Decode(s))
		}
		s = append(s, *p)
		pp := uintptr(unsafe.Pointer(p))
		pp += 2 // sizeof(uint16)
		p = (*uint16)(unsafe.Pointer(pp))
	}
}
