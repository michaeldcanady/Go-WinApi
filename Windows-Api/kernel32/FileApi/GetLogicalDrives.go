package fileapi

import "errors"

//GetLogicalDrives returns a list of all logical drives on the host machine
func GetLogicalDrives() ([]string, error) {
	ret, _, _ := procGetLogicalDrives.Call()
	if ret == 0 {
		return []string{}, errors.New("no drives found")
	}
	return bitsToDrives(uint32(ret)), nil
}
