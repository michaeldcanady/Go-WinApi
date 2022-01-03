package fileapi

//DeleteFileW Deletes an existing file.
//
//To perform this operation as a transacted operation, use the DeleteFileTransacted function.
func DeleteFileW(fileName string) error {

	ret, _, err := procDeleteFileW.Call(UintptrFromString(fileName))

	if ret == 0 {
		return err
	}

	return nil
}
