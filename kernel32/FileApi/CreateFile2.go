package fileapi

func CreateFile2(fileName string, dwDesiredAccess, dwShareMode, dwCreationDisposition DWORD) (HANDLE, error) {

	ret, _, err := procCreateFile2.Call(
		UintptrFromString(fileName),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(dwCreationDisposition),
		0,
	)

	if ret == 18446744073709551615 {
		return HANDLE(ret), err
	}

	return HANDLE(ret), nil
}
