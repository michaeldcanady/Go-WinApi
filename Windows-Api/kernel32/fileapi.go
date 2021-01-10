package kernel32

import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	findFirstVolumeWProc                  = kernel32.NewProc("FindFirstVolumeW")
	findNextVolumeWProc                   = kernel32.NewProc("FindNextVolumeW")
	findVolumeCloseProc                   = kernel32.NewProc("FindVolumeClose")
	getVolumePathNamesForVolumeNameWProc  = kernel32.NewProc("GetVolumePathNamesForVolumeNameW")
	procGetVolumeNameForVolumeMountPointW = kernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	getDriveTypeWProc                     = kernel32.NewProc("GetDriveTypeW")
	procGetVolumeInformationW             = kernel32.NewProc("GetVolumeInformationW")
	procGetLogicalDrives                  = kernel32.NewProc("GetLogicalDrives")
	procDeleteVolumeMountPointW           = kernel32.NewProc("DeleteVolumeMountPointW")
	procGetDriveTypeW                     = kernel32.NewProc("GetDriveTypeW")

	errUnknownDriveType = errors.New("unknown drive type")
	errNoRootDir        = errors.New("invalid root drive path")

	driveTypeErrors = [...]error{
		0: errUnknownDriveType,
		1: errNoRootDir,
	}
)

type fixedDriveVolume struct {
	volName          string
	mountedPathnames []string
}

type fixedVolumeMounts struct {
	volume string
	mounts []string
}

type Volume struct {
	PathName                 string
	VolumeLabel              string
	nVolumeNameSize          uint32
	SerialNumber             uint32
	lpMaximumComponentLength uint32
	SystemFlags              uint32
	FileSystem               string
	nFileSystemNameSize      uint32
}

func New(lpRootPathName string, lpVolumeNameBuffer []uint16, nVolumeNameSize, lpVolumeSerialNumber, lpMaximumComponentLength, lpFileSystemFlags uint32, lpFileSystemNameBuffer []uint16, nFileSystemNameSize uint32) Volume {
	label := syscall.UTF16ToString(lpVolumeNameBuffer)
	if label == "" {
		label = "Local Disk"
	}
	return Volume{
		PathName:                 lpRootPathName,
		VolumeLabel:              label,
		nVolumeNameSize:          nVolumeNameSize,
		SerialNumber:             lpVolumeSerialNumber,
		lpMaximumComponentLength: lpMaximumComponentLength,
		SystemFlags:              lpFileSystemFlags,
		FileSystem:               syscall.UTF16ToString(lpFileSystemNameBuffer),
		nFileSystemNameSize:      nFileSystemNameSize,
	}
}

func GetDriveTypeW(PathName string) (string, error) {

	vmpp, err := syscall.UTF16PtrFromString(PathName)
	if err != nil {
		fmt.Println(err)
	}

	ret, _, err := procGetDriveTypeW.Call(uintptr(unsafe.Pointer(vmpp)))
	if ret == 0 {
		fmt.Println(err)
	}

	switch ret {
	case DRIVE_UNKNOWN:
		return "unknown type", nil
	case DRIVE_NO_ROOT_DIR:
		return "no root dir", nil
	case DRIVE_REMOVABLE:
		return "removable", nil
	case DRIVE_FIXED:
		return "fixed", nil
	case DRIVE_REMOTE:
		return "remote", nil
	case DRIVE_CDROM:
		return "cdrom", nil
	case DRIVE_RAMDISK:
		return "ramdisk", nil
	}
	return "", nil
}

func DeleteVolumeMountPointW(volumeMountPoint string) {

	vmpp, err := syscall.UTF16PtrFromString(volumeMountPoint)
	if err != nil {
		fmt.Println(err)
	}

	ret, _, err := procDeleteVolumeMountPointW.Call(uintptr(unsafe.Pointer(vmpp)))
	if ret == 0 {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func GetVolumeInformationW(rootPathName string) Volume {
	if !strings.HasSuffix(rootPathName, "\\") {
		rootPathName = rootPathName + "\\"
	}

	var VolumeNameBuffer = make([]uint16, syscall.MAX_PATH+1)
	var nVolumeNameSize = uint32(len(VolumeNameBuffer))
	var VolumeSerialNumber uint32
	var MaximumComponentLength uint32
	var FileSystemFlags uint32
	var FileSystemNameBuffer = make([]uint16, 255)
	var nFileSystemNameSize uint32 = syscall.MAX_PATH + 1

	_, _, err := procGetVolumeInformationW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(rootPathName))),
		uintptr(unsafe.Pointer(&VolumeNameBuffer[0])),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(&VolumeSerialNumber)),
		uintptr(unsafe.Pointer(&MaximumComponentLength)),
		uintptr(unsafe.Pointer(&FileSystemFlags)),
		uintptr(unsafe.Pointer(&FileSystemNameBuffer[0])),
		uintptr(nFileSystemNameSize),
		0)
	if err != nil {

	}
	return New(rootPathName, VolumeNameBuffer, nVolumeNameSize, VolumeSerialNumber, MaximumComponentLength, FileSystemFlags, FileSystemNameBuffer, nFileSystemNameSize)
}

func GetLogicalDrives() ([]string, error) {
	ret, _, _ := procGetLogicalDrives.Call()
	if ret == 0 {
		return []string{}, errors.New("No drives found.")
	}
	return bitsToDrives(uint32(ret)), nil
}

func GetVolumeNameForVolumeMountPointW(volumeMountPoint string) (string, error) {

	if len(volumeMountPoint) == 0 {
		return "", syscall.EINVAL
	}
	if !strings.HasSuffix(volumeMountPoint, "\\") {
		volumeMountPoint = volumeMountPoint + "\\"
	}

	vmpp, err := syscall.UTF16PtrFromString(volumeMountPoint)
	if err != nil {
		return "", err
	}

	var vnBuffer [MaxVolumeNameLength]uint16
	p0 := &vnBuffer[0]

	re, _, err := procGetVolumeNameForVolumeMountPointW.Call(
		uintptr(unsafe.Pointer(vmpp)),
		uintptr(unsafe.Pointer(p0)),
		uintptr(MaxVolumeNameLength),
	)
	if re == 0 {
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(vnBuffer[:]), nil
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
