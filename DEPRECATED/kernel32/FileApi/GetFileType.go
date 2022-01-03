package fileapi

var (
	fileType = map[int64]string{
		0x0001: "FILE_TYPE_DISK",
		0x0002: "FILE_TYPE_CHAR",
		0x0003: "FILE_TYPE_PIPE",
		0x8000: "FILE_TYPE_REMOTE",
	}
)

func GetFileType(hFile HANDLE) (string, error) {

	ret, _, err := procGetFileType.Call(uintptr(hFile))

	if ret == 0 {
		return "FILE_TYPE_UNKNOWN", err
	}

	return SeperateFlags(uint32(ret), fileType)[0], nil
}
