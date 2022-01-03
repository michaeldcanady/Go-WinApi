package GoWinApi

import (
	"time"
	"unsafe"
)

var (
	fileTimeToSystemTimeProc = kernel32.NewProc("FileTimeToSystemTime")
	SystemTimeToFileTimeProc = kernel32.NewProc("SystemTimeToFileTime")
)

//FileTime Contains a 64-bit value representing the number of 100-nanosecond intervals since January 1, 1601 (UTC).
type FileTime struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

func (F FileTime) ToUintPtr() uintptr {
	return uintptr(unsafe.Pointer(&F))
}

//SystemTime Specifies a date and time, using individual members for the month, day, year, weekday, hour, minute, second, and millisecond.
//The time is either in coordinated universal time (UTC) or local time, depending on the function that is being called.
type SystemTime struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

func (S SystemTime) ToDate() time.Time {
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

func (S SystemTime) ToUintPtr() uintptr {
	return uintptr(unsafe.Pointer(&S))
}

//FileTimeToSystemTime Converts a file time to system time format. System time is based on Coordinated Universal Time (UTC).
func FileTimeToSystemTime(lpFileTime FileTime) (time.Time, error) {

	var lpSystemTime SystemTime

	ret, _, err := fileTimeToSystemTimeProc.Call(
		uintptr(unsafe.Pointer(&lpFileTime)),
		uintptr(unsafe.Pointer(&lpSystemTime)),
	)

	if ret == 0 {
		return lpSystemTime.ToDate(), err
	}

	return lpSystemTime.ToDate(), nil
}

//SystemTimeToFileTime Converts a system time to file time format. System time is based on Coordinated Universal Time (UTC).
func SystemTimeToFileTime() {

}
