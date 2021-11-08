package fileapi

func FindClose(hFindFile HANDLE) error {
	
	ret, _, err := procFindClose.Call(uintptr(hFindFile))

	if ret == 0 {
		return err
	}

	return nil
}
