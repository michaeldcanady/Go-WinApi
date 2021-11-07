package fileapi

var areFileApisANSIProc = kernel32.NewProc("AreFileApisANSI")

//AreFileApisANSI If the set of file I/O functions is using the ANSI code page, the return value is nonzero.
//If the set of file I/O functions is using the OEM code page, the return value is zero.
func AreFileApisANSI() bool {
	ret,_,_ := areFileApisANSIProc.Call()

	return ret != 0
}
