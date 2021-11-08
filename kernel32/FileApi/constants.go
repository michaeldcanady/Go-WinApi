package fileapi

import "syscall"

const (
	guidBufLen          = syscall.MAX_PATH + 1
	MaxVolumeNameLength = 50
	driveUnknown        = iota
	driveNoRootDir

	driveRemovable
	driveFixed
	driveRemote
	driveCDROM
	driveRamdisk

	driveLastKnownType = driveRamdisk

	DRIVE_UNKNOWN     = 0
	DRIVE_NO_ROOT_DIR = 1
	DRIVE_REMOVABLE   = 2
	DRIVE_FIXED       = 3
	DRIVE_REMOTE      = 4
	DRIVE_CDROM       = 5
	DRIVE_RAMDISK     = 6
)
