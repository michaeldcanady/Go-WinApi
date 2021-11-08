package fileapi

func DeleteVolumeMountPointW(volumeMountPoint string) error {

	ret, _, err := procDeleteVolumeMountPointW.Call(UintptrFromString(volumeMountPoint))
	if ret == 0 {
		return err
	}
	
	return nil
}
