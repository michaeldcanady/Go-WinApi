package fileapi

func GetFileSizeEx(hFile HANDLE) (int64, error) {

	ret, _, err := procGetFileSizeEx.Call(
		uintptr(hFile),
		0,
	)
	if ret == 0 {
		return 0, err
	}

	return int64(ret), nil
}
