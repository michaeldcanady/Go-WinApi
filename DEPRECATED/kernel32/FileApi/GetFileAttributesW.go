package fileapi

func GetFileAttributesW(fileName string) ([]string, error) {
	ret, _, err := procGetFileAttributesW.Call(UintptrFromString(fileName))

	if ret == 0xFFFFFFFF {
		return []string{}, err
	}

	return SeperateFlags(uint32(ret), dwFileAttributeFlags), err
}
