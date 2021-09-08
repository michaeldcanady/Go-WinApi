package fileapi

func findVolumeClose(handle uintptr) error {
	ok, _, err := findVolumeCloseProc.Call(handle)
	if ok == 0 {
		return err
	}

	return nil
}
