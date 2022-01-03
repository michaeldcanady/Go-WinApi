package fileapi

func GetFileSize(hFile HANDLE) (int64, error) {

	ret, _, err := procGetFileSize.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0xFFFFFFFF {
		return 0, err
	}

	return int64(ret), nil
}
