package fileapi

//FindClose Closes a file search handle opened by the FindFirstFile, FindFirstFileEx, FindFirstFileNameW, FindFirstFileNameTransactedW, FindFirstFileTransacted,
//FindFirstStreamTransactedW, or FindFirstStreamW functions.
func FindClose(hFindFile HANDLE) error {

	ret, _, err := procFindClose.Call(hFindFile.ToUintPtr())

	if ret == 0 {
		return err
	}

	return nil
}
