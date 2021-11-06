package sysinfoapi

import (
	"fmt"
	"unsafe"
)

var (
	procGetSystemFirmwareTable = kernel32.NewProc("GetSystemFirmwareTable")
)

type FirmwareTableId int64

const (
	FACP FirmwareTableId = 1178682192
	PCAF FirmwareTableId = 1346584902
	MSDM FirmwareTableId = 1296323405
)

func GetSystemFirmwareTable(firmwareTableProviderSignature FirmwareTable, FirmwareTableID FirmwareTableId) {

	ret, _, err := procGetSystemFirmwareTable.Call(
		uintptr(firmwareTableProviderSignature),
		uintptr(FirmwareTableID),
		uintptr(unsafe.Pointer(&make([]uint8, 1)[0])),
		uintptr(1),
	)

	if ret == 0 {
		fmt.Println(err)
	}

	pFirmwareTableEnumBuffer := make([]uint8, ret)

	ret, _, err = procGetSystemFirmwareTable.Call(
		uintptr(firmwareTableProviderSignature),
		uintptr(FirmwareTableID),
		uintptr(unsafe.Pointer(&pFirmwareTableEnumBuffer[0])),
		uintptr(ret),
	)

	if ret == 0 {
		fmt.Println(err)
	}

	//var tables []string

	fmt.Println(pFirmwareTableEnumBuffer[56:])
}
