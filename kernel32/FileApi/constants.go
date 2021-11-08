package fileapi

import "syscall"

const (
	MAX_PATH            = syscall.MAX_PATH
	guidBufLen          = MAX_PATH + 1
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

	//Possible Security Attributes
	FILE_ATTRIBUTE_ARCHIVE   SecurityAttribute = 32
	FILE_ATTRIBUTE_ENCRYPTED SecurityAttribute = 16384
	FILE_ATTRIBUTE_HIDDEN    SecurityAttribute = 2
	FILE_ATTRIBUTE_NORMAL    SecurityAttribute = 128 //Only Used Alone
	FILE_ATTRIBUTE_OFFLINE   SecurityAttribute = 4096
	FILE_ATTRIBUTE_READONLY  SecurityAttribute = 1
	FILE_ATTRIBUTE_SYSTEM    SecurityAttribute = 4
	FILE_ATTRIBUTE_TEMPORARY SecurityAttribute = 256

	//dwAdditionalFlags
	FIND_FIRST_EX_CASE_SENSITIVE       = 1
	FIND_FIRST_EX_LARGE_FETCH          = 2
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY = 4
	
	//FINDEX_INFO_LEVELS
	FindExInfoStandard     = 0
	FindExInfoBasic        = 1
	FindExInfoMaxInfoLevel = 2
	
	//FINDEX_SEARCH_OPS
	FindExSearchNameMatch          = 0
	FindExSearchLimitToDirectories = 1
	FindExSearchLimitToDevices     = 2
	FindExSearchMaxSearchOp        = 3
	
	MAXDWORD = 4294967295

	FILE_SUPPORTS_USN_JOURNAL uint32 = 0x02000000

	//File type values
	FILE_TYPE_UNKNOWN = 0x0000
	FILE_TYPE_DISK    = 0x0001
	FILE_TYPE_CHAR    = 0x0002
	FILE_TYPE_PIPE    = 0x0003
	FILE_TYPE_REMOTE  = 0x8000
)
