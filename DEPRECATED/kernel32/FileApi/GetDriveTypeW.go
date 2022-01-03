package fileapi

func GetDriveTypeW(PathName string) (string, error) {

	ret, _, err := procGetDriveTypeW.Call(UintptrFromString(PathName))

	return DriveType(ret).String(), err
}
