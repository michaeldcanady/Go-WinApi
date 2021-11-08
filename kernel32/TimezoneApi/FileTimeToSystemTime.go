package timezoneapi

import (
	"time"
	"unsafe"
)

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

func (F FILETIME) toUTF16Ptr() uintptr {
	return uintptr(unsafe.Pointer(&F))
}

type SYSTEMTIME struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

func (S SYSTEMTIME) ToDate() time.Time {
	return time.Date(
		int(S.wYear),
		time.Month(S.wMonth),
		int(S.wDay),
		int(S.wHour),
		int(S.wMinute),
		int(S.wSecond),
		int(S.wMilliseconds*1000),
		time.UTC,
	)
}

func FileTimeToSystemTime(lpFileTime FILETIME) (time.Time, error) {

	var lpSystemTime SYSTEMTIME

	ret, _, err := fileTimeToSystemTimeProc.Call(
		uintptr(unsafe.Pointer(&lpFileTime)),
		uintptr(unsafe.Pointer(&lpSystemTime)),
	)

	if ret == 0 {
		return lpSystemTime.ToDate(), err
	}

	return lpSystemTime.ToDate(), nil
}
