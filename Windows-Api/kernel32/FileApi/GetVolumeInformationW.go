package fileapi

import (
	"strings"
	"syscall"
	"strconv"
	"unsafe"
	"fmt"
)

const (
	FILE_CASE_PRESERVED_NAMES         = 0x00000002
	FILE_CASE_SENSITIVE_SEARCH        = 0x00000001
	FILE_DAX_VOLUME                   = 0x20000000
	FILE_FILE_COMPRESSION             = 0x00000010
	FILE_NAMED_STREAMS                = 0x00040000
	FILE_PERSISTENT_ACLS              = 0x00000008
	FILE_READ_ONLY_VOLUME             = 0x00080000
	FILE_SEQUENTIAL_WRITE_ONCE        = 0x00100000
	FILE_SUPPORTS_ENCRYPTION          = 0x00020000
	FILE_SUPPORTS_EXTENDED_ATTRIBUTES = 0x00800000
	FILE_SUPPORTS_HARD_LINKS          = 0x00400000
	FILE_SUPPORTS_OBJECT_IDS          = 0x00010000
	FILE_SUPPORTS_OPEN_BY_FILE_ID     = 0x01000000
	FILE_SUPPORTS_REPARSE_POINTS      = 0x00000080
	FILE_SUPPORTS_SPARSE_FILES        = 0x00000040
	FILE_SUPPORTS_TRANSACTIONS        = 0x00200000
	FILE_SUPPORTS_USN_JOURNAL         = 0x02000000
	FILE_UNICODE_ON_DISK              = 0x00000004
	FILE_VOLUME_IS_COMPRESSED         = 0x00008000
	FILE_VOLUME_QUOTAS                = 0x00000020
	FILE_SUPPORTS_BLOCK_REFCOUNTING   = 0x08000000
)

func GetVolumeInformationW(rootPathName string) (Volume, error) {
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

	//Converts rootPathName to UTF16 Pointer requred by procGetVolumeInformationW
	pointer, err := syscall.UTF16PtrFromString(rootPathName)
	if err != nil {
		return Volume{}, err
	}

	ret, _, err := procGetVolumeInformationW.Call(
		uintptr(unsafe.Pointer(pointer)),
		uintptr(unsafe.Pointer(&VolumeNameBuffer[0])),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(&VolumeSerialNumber)),
		uintptr(unsafe.Pointer(&MaximumComponentLength)),
		uintptr(unsafe.Pointer(&FileSystemFlags)),
		uintptr(unsafe.Pointer(&FileSystemNameBuffer[0])),
		uintptr(nFileSystemNameSize),
		0)
	// If GetVolumeInformationW Call returns an error
	if ret != 1 {
		return Volume{}, err
	}

	intFlags := seperateFlags(FileSystemFlags)

	flags := identifyFlags(intFlags)

	//Returns new volume
	return newVolume(rootPathName, VolumeNameBuffer, nVolumeNameSize, VolumeSerialNumber, MaximumComponentLength, flags, FileSystemNameBuffer, nFileSystemNameSize), nil
}

//identifyFlags Determines flags on drive based off consts
func identifyFlags(flags []int64) (identifiedFlags []string) {
	for _, flag := range flags {
		switch flag {
		case FILE_CASE_PRESERVED_NAMES:
			identifiedFlags = append(identifiedFlags, "FILE_CASE_PRESERVED_NAMES")
		case FILE_CASE_SENSITIVE_SEARCH:
			identifiedFlags = append(identifiedFlags, "FILE_CASE_SENSITIVE_SEARCH")
		case FILE_DAX_VOLUME:
			identifiedFlags = append(identifiedFlags, "FILE_DAX_VOLUME")
		case FILE_FILE_COMPRESSION:
			identifiedFlags = append(identifiedFlags, "FILE_FILE_COMPRESSION")
		case FILE_NAMED_STREAMS:
			identifiedFlags = append(identifiedFlags, "FILE_NAMED_STREAMS")
		case FILE_PERSISTENT_ACLS:
			identifiedFlags = append(identifiedFlags, "FILE_PERSISTENT_ACLS")
		case FILE_READ_ONLY_VOLUME:
			identifiedFlags = append(identifiedFlags, "FILE_READ_ONLY_VOLUME")
		case FILE_SEQUENTIAL_WRITE_ONCE:
			identifiedFlags = append(identifiedFlags, "FILE_SEQUENTIAL_WRITE_ONCE")
		case FILE_SUPPORTS_ENCRYPTION:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_ENCRYPTION")
		case FILE_SUPPORTS_EXTENDED_ATTRIBUTES:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_EXTENDED_ATTRIBUTES")
		case FILE_SUPPORTS_HARD_LINKS:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_HARD_LINKS")
		case FILE_SUPPORTS_OBJECT_IDS:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_OBJECT_IDS")
		case FILE_SUPPORTS_OPEN_BY_FILE_ID:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_OPEN_BY_FILE_ID")
		case FILE_SUPPORTS_REPARSE_POINTS:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_REPARSE_POINTS")
		case FILE_SUPPORTS_SPARSE_FILES:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_SPARSE_FILES")
		case FILE_SUPPORTS_TRANSACTIONS:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_TRANSACTIONS")
		case FILE_SUPPORTS_USN_JOURNAL:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_USN_JOURNAL")
		case FILE_UNICODE_ON_DISK:
			identifiedFlags = append(identifiedFlags, "FILE_UNICODE_ON_DISK")
		case FILE_VOLUME_IS_COMPRESSED:
			identifiedFlags = append(identifiedFlags, "FILE_VOLUME_IS_COMPRESSED")
		case FILE_VOLUME_QUOTAS:
			identifiedFlags = append(identifiedFlags, "FILE_VOLUME_QUOTAS")
		case FILE_SUPPORTS_BLOCK_REFCOUNTING:
			identifiedFlags = append(identifiedFlags, "FILE_SUPPORTS_BLOCK_REFCOUNTING")

		}
	}
	return
}

//parseBinToHex converts binary to hex
func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
