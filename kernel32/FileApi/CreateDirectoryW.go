package fileapi

func CreateDirectoryW(pathName string, SecurityAttributes SecurityAttribute) error {

	r, _, err := procCreateDirectoryW.Call(
		UintptrFromString(pathName), // [in] LPCTSTR
		uintptr(SecurityAttributes), // [in] LPSECURITY_ATT...
	)

	if r == 0 {
		return err
	}

	return nil
}
