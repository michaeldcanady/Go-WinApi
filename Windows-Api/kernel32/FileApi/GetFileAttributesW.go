package fileapi

var getFileAttributesWProc = kernel32.NewProc("GetFileAttributesW")

func GetFileAttributesW(lpFileName string) ([]string, error) {
	ret, _, err := getFileAttributesWProc.Call(UintptrFromString(&lpFileName))

	if ret == 0xFFFFFFFF {
		return []string{}, err
	}

	return seperateFlags(uint32(ret), dwFileAttributeFlags), err
}
