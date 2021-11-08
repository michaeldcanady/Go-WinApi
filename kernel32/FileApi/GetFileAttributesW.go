package fileapi

func GetFileAttributesW(lpFileName string) ([]string, error) {
	ret, _, err := getFileAttributesWProc.Call(UintptrFromString(&lpFileName))

	if ret == 0xFFFFFFFF {
		return []string{}, err
	}

	return SeperateFlags(uint32(ret), dwFileAttributeFlags), err
}
