package fileapi

func findVolumeClose(handle HANDLE) error {
	ok, _, err := procFindVolumeClose.Call(uintptr(handle))

	if ok == 0 {
		return err
	}

	return nil
}
