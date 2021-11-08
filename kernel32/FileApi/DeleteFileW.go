package fileapi

func DeleteFileW(fileName string) error {
	
	ret, _, err := procDeleteFileW.Call(UintptrFromString(fileName))

	if ret == 0 {
		return err
	}

	return nil
}
