package fileapi

func CreateDirectoryW(pathName string, SecurityAttributes SecurityAttribute) error {

	r, _, err := procCreateDirectoryW.Call(
		UintptrFromString(pathName),
		uintptr(SecurityAttributes),
	)

	if r == 0 {
		return err
	}

	return nil
}
