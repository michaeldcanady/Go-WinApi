package fileapi

import "syscall"

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
	SystemFlags              []string
	FileSystem               string
	nFileSystemNameSize      uint32
}

func newVolume(lpRootPathName string, lpVolumeNameBuffer []uint16, nVolumeNameSize, lpVolumeSerialNumber, lpMaximumComponentLength uint32, lpFileSystemFlags uint32, lpFileSystemNameBuffer []uint16, nFileSystemNameSize uint32) Volume {
	label := syscall.UTF16ToString(lpVolumeNameBuffer)
	return Volume{
		PathName:                 lpRootPathName,
		VolumeLabel:              label,
		nVolumeNameSize:          nVolumeNameSize,
		SerialNumber:             lpVolumeSerialNumber,
		lpMaximumComponentLength: lpMaximumComponentLength,
		SystemFlags:              seperateFlags(lpFileSystemFlags, volumeFlags),
		FileSystem:               syscall.UTF16ToString(lpFileSystemNameBuffer),
		nFileSystemNameSize:      nFileSystemNameSize,
	}
}
