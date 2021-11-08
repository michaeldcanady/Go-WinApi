package fileapi

func CreateFile2(fileName string, dwDesiredAccess, dwShareMode, dwCreationDisposition uint32) (handle HANDLE, err error) {

	ret, _, err := procCreateFile2.Call(
		UintptrFromString(fileName),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(dwCreationDisposition),
		0,
	)

	handle = HANDLE(ret)

	if handle == INVALID_HANDLE_VALUE {
		return handle, err
	}

	return handle, nil
}
