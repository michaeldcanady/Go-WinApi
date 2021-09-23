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
