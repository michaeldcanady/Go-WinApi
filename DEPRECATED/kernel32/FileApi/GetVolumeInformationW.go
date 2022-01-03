package fileapi

import (
	"strings"
	"syscall"
	"unsafe"
)

type VolumeFlags int64

const (
	FILE_CASE_PRESERVED_NAMES         VolumeFlags = 0x00000002
	FILE_CASE_SENSITIVE_SEARCH        VolumeFlags = 0x00000001
	FILE_DAX_VOLUME                   VolumeFlags = 0x20000000
	FILE_FILE_COMPRESSION             VolumeFlags = 0x00000010
	FILE_NAMED_STREAMS                VolumeFlags = 0x00040000
	FILE_PERSISTENT_ACLS              VolumeFlags = 0x00000008
	FILE_READ_ONLY_VOLUME             VolumeFlags = 0x00080000
	FILE_SEQUENTIAL_WRITE_ONCE        VolumeFlags = 0x00100000
	FILE_SUPPORTS_ENCRYPTION          VolumeFlags = 0x00020000
	FILE_SUPPORTS_EXTENDED_ATTRIBUTES VolumeFlags = 0x00800000
	FILE_SUPPORTS_HARD_LINKS          VolumeFlags = 0x00400000
	FILE_SUPPORTS_OBJECT_IDS          VolumeFlags = 0x00010000
	FILE_SUPPORTS_OPEN_BY_FILE_ID     VolumeFlags = 0x01000000
	FILE_SUPPORTS_REPARSE_POINTS      VolumeFlags = 0x00000080
	FILE_SUPPORTS_SPARSE_FILES        VolumeFlags = 0x00000040
	FILE_SUPPORTS_TRANSACTIONS        VolumeFlags = 0x00200000
	FILE_SUPPORTS_USN_JOURNAL         VolumeFlags = 0x02000000
	FILE_UNICODE_ON_DISK              VolumeFlags = 0x00000004
	FILE_VOLUME_IS_COMPRESSED         VolumeFlags = 0x00008000
	FILE_VOLUME_QUOTAS                VolumeFlags = 0x00000020
	FILE_SUPPORTS_BLOCK_REFCOUNTING   VolumeFlags = 0x08000000
)

var (
	volumeFlags = map[int64]string{
		0x00000002: "FILE_CASE_PRESERVED_NAMES",
		0x00000001: "FILE_CASE_SENSITIVE_SEARCH",
		0x20000000: "FILE_DAX_VOLUME",
		0x00000010: "FILE_FILE_COMPRESSION",
		0x00040000: "FILE_NAMED_STREAMS",
		0x00000008: "FILE_PERSISTENT_ACLS",
		0x00080000: "FILE_READ_ONLY_VOLUME",
		0x00100000: "FILE_SEQUENTIAL_WRITE_ONCE",
		0x00020000: "FILE_SUPPORTS_ENCRYPTION",
		0x00800000: "FILE_SUPPORTS_EXTENDED_ATTRIBUTES",
		0x00400000: "FILE_SUPPORTS_HARD_LINKS",
		0x00010000: "FILE_SUPPORTS_OBJECT_IDS",
		0x01000000: "FILE_SUPPORTS_OPEN_BY_FILE_ID",
		0x00000080: "FILE_SUPPORTS_REPARSE_POINTS",
		0x00000040: "FILE_SUPPORTS_SPARSE_FILES",
		0x00200000: "FILE_SUPPORTS_TRANSACTIONS",
		0x02000000: "FILE_SUPPORTS_USN_JOURNAL",
		0x00000004: "FILE_UNICODE_ON_DISK",
		0x00008000: "FILE_VOLUME_IS_COMPRESSED",
		0x00000020: "FILE_VOLUME_QUOTAS",
		0x08000000: "FILE_SUPPORTS_BLOCK_REFCOUNTING",
	}
)

func GetVolumeInformationW(rootPathName string) (volume Volume, err error) {
	if !strings.HasSuffix(rootPathName, "\\") {
		rootPathName = rootPathName + "\\"
	}

	var (
		VolumeSerialNumber     uint32
		MaximumComponentLength uint32
		FileSystemFlags        uint32
		VolumeNameBuffer              = make([]uint16, syscall.MAX_PATH+1)
		nVolumeNameSize               = uint32(len(VolumeNameBuffer))
		FileSystemNameBuffer          = make([]uint16, 255)
		nFileSystemNameSize    uint32 = syscall.MAX_PATH + 1
	)

	ret, _, err := procGetVolumeInformationW.Call(
		UintptrFromString(rootPathName),
		uintptr(unsafe.Pointer(&VolumeNameBuffer[0])),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(&VolumeSerialNumber)),
		uintptr(unsafe.Pointer(&MaximumComponentLength)),
		uintptr(unsafe.Pointer(&FileSystemFlags)),
		uintptr(unsafe.Pointer(&FileSystemNameBuffer[0])),
		uintptr(nFileSystemNameSize),
		0)
	// If GetVolumeInformationW Call returns an error
	if ret == 0 {
		return
	}

	volume = newVolume(rootPathName, VolumeNameBuffer, nVolumeNameSize, VolumeSerialNumber, MaximumComponentLength, FileSystemFlags, nFileSystemNameSize, FileSystemNameBuffer)
	//Returns new volume
	return volume, nil
}