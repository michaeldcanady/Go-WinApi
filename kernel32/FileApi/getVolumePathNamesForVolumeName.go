package fileapi

import (
	"syscall"
	"unsafe"
)

func GetVolumePathNamesForVolumeName(volName []string) ([]string, error) {
	const (
		errorMoreData = 234
		NUL           = 0x0000
	)

	var (
		pathNamesLen uint32
		pathNames    []uint16
	)

	pathNamesLen = 2
	for {
		pathNames = make([]uint16, pathNamesLen)
		pathNamesLen *= 2

		lpszVolumeName, err := syscall.UTF16PtrFromString(volName[0])

		if err != nil {
			return []string{}, err
		}

		rc, _, err := getVolumePathNamesForVolumeNameWProc.Call(
			lpszVolumeName,
			uintptr(unsafe.Pointer(&pathNames[0])),
			uintptr(pathNamesLen),
			uintptr(unsafe.Pointer(&pathNamesLen)),
		)

		if rc == 0 {
			if err.(syscall.Errno) == errorMoreData {
				continue
			}

			return nil, err
		}

		pathNames = pathNames[:pathNamesLen]
		break
	}

	var out [][]uint16
	i := 0
	for j, c := range pathNames {
		if c == NUL && i < j {
			out = append(out, pathNames[i:j+1])
			i = j + 1
		}
	}

	return LPSTRsToStrings(out), nil
}