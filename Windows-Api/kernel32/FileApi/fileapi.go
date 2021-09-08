package fileapi

import (
	"errors"
	"syscall"
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
