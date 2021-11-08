package fileapi

import "fmt"

func GetDriveTypeW(PathName string) (string, error) {

	ret, _, err := procGetDriveTypeW.Call(UintptrFromString(PathName))
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
