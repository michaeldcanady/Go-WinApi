package fileapi

var (
	procDeleteFileW = kernel32.NewProc("DeleteFileW")
)

func DeleteFileW(FileName string) error {
	ret, _, err := procDeleteFileW.Call(UintptrFromString(&FileName))

	if ret == 0 {
		return err
	}

	return nil
}
