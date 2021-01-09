package Winapi

import (
	"errors"
	"log"
	"strings"
	"syscall"
	"unsafe"
)

const (
  guidBufLen   = syscall.MAX_PATH + 1
	driveUnknown = iota
	driveNoRootDir

	driveRemovable
	driveFixed
	driveRemote
	driveCDROM
	driveRamdisk

	driveLastKnownType = driveRamdisk
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	findFirstVolumeWProc = kernel32.NewProc("FindFirstVolumeW")
	findNextVolumeWProc  = kernel32.NewProc("FindNextVolumeW")
	findVolumeCloseProc  = kernel32.NewProc("FindVolumeClose")

	getVolumePathNamesForVolumeNameWProc = kernel32.NewProc("GetVolumePathNamesForVolumeNameW")

	getDriveTypeWProc = kernel32.NewProc("GetDriveTypeW")
)

type fixedDriveVolume struct {
	volName          string
	mountedPathnames []string
}

type fixedVolumeMounts struct {
	volume string
	mounts []string
}

func main() {
	mounts, err := getFixedDriveMounts()
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range mounts {
		log.Println("volume:", m.volume,
			"mounts:", strings.Join(m.mounts, ", "))
	}
}

func findFirstVolume() (uintptr, []uint16, error) {
	const invalidHandleValue = ^uintptr(0)

	guid := make([]uint16, guidBufLen)

	handle, _, err := findFirstVolumeWProc.Call(
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if handle == invalidHandleValue {
		return invalidHandleValue, nil, err
	}

	return handle, guid, nil
}

func findNextVolume(handle uintptr) ([]uint16, bool, error) {
	const noMoreFiles = 18

	guid := make([]uint16, guidBufLen)

	rc, _, err := findNextVolumeWProc.Call(
		handle,
		uintptr(unsafe.Pointer(&guid[0])),
		uintptr(guidBufLen*2),
	)

	if rc == 1 {
		return guid, true, nil
	}

	if err.(syscall.Errno) == noMoreFiles {
		return nil, false, nil
	}
	return nil, false, err
}

func findVolumeClose(handle uintptr) error {
	ok, _, err := findVolumeCloseProc.Call(handle)
	if ok == 0 {
		return err
	}

	return nil
}

func getVolumePathNamesForVolumeName(volName []uint16) ([][]uint16, error) {
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

		rc, _, err := getVolumePathNamesForVolumeNameWProc.Call(
			uintptr(unsafe.Pointer(&volName[0])),
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
	return out, nil
}

func getDriveType(rootPathName []uint16) (int, error) {
	rc, _, _ := getDriveTypeWProc.Call(
		uintptr(unsafe.Pointer(&rootPathName[0])),
	)

	dt := int(rc)

	if dt == driveUnknown || dt == driveNoRootDir {
		return -1, driveTypeErrors[dt]
	}

	return dt, nil
}

func getFixedDriveMounts() ([]fixedVolumeMounts, error) {
	var out []fixedVolumeMounts

	err := enumVolumes(func(guid []uint16) error {
		mounts, err := maybeGetFixedVolumeMounts(guid)
		if err != nil {
			return err
		}
		if len(mounts) > 0 {
			out = append(out, fixedVolumeMounts{
				volume: syscall.UTF16ToString(guid),
				mounts: LPSTRsToStrings(mounts),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

func enumVolumes(handleVolume func(guid []uint16) error) error {
	handle, guid, err := findFirstVolume()
	if err != nil {
		return err
	}
	defer func() {
		err = findVolumeClose(handle)
	}()

	if err := handleVolume(guid); err != nil {
		return err
	}

	for {
		guid, more, err := findNextVolume(handle)
		if err != nil {
			return err
		}

		if !more {
			break
		}

		if err := handleVolume(guid); err != nil {
			return err
		}
	}

	return nil
}

func maybeGetFixedVolumeMounts(guid []uint16) ([][]uint16, error) {
	paths, err := getVolumePathNamesForVolumeName(guid)
	if err != nil {
		return nil, err
	}

	if len(paths) == 0 {
		return nil, nil
	}

	var lastErr error
	for _, path := range paths {
		dt, err := getDriveType(path)
		if err == nil {
			if dt == driveFixed {
				return paths, nil
			}
			return nil, nil
		}
		lastErr = err
	}

	return nil, lastErr
}
