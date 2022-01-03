package GoWinApi

import (
	"fmt"
	"unsafe"
)

type (
	FirmwareTable   int64
	FirmwareTableId int64
)

var (
	procEnumSystemFirmwareTables = kernel32.NewProc("EnumSystemFirmwareTables")
	procGetSystemFirmwareTable   = kernel32.NewProc("GetSystemFirmwareTable")
)

const (
	//ACPI The ACPI firmware table provider.
	ACPI FirmwareTable = 1094930505
	//FIRM The raw firmware table provider. Not supported for UEFI systems; use 'RSMB' instead.
	FIRM FirmwareTable = 1179210317
	//RSMB The raw SMBIOS firmware table provider.
	RSMB FirmwareTable = 1381190978

	FACP FirmwareTableId = 1178682192
	PCAF FirmwareTableId = 1346584902
	MSDM FirmwareTableId = 1296323405
)

//EnumSystemFirmwareTables Enumerates all system firmware tables of the specified type.
func EnumSystemFirmwareTables(firmwareTableProviderSignature FirmwareTable) (tables []string, err error) {

	var ret uintptr = 1
	pFirmwareTableEnumBuffer := make([]uint8, ret)

	for i := 0; i < 2; i++ {
		ret, _, err = procEnumSystemFirmwareTables.Call(
			uintptr(firmwareTableProviderSignature),
			uintptr(unsafe.Pointer(&pFirmwareTableEnumBuffer[0])),
			uintptr(ret),
		)

		if ret == 0 {
			return
		}

	}

	for i := 0; i < len(pFirmwareTableEnumBuffer); i += 4 {
		tables = append(tables, string(pFirmwareTableEnumBuffer[i:i+4]))
	}

	return
}

//GetSystemFirmwareTable Retrieves the specified firmware table from the firmware table provider.
func GetSystemFirmwareTable(firmwareTableProviderSignature FirmwareTable, FirmwareTableID FirmwareTableId) {

	ret := 0
	var pFirmwareTableEnumBuffer []uint8

	for i := 0; i < 2; i++ {
		pFirmwareTableEnumBuffer = make([]uint8, ret)

		ret, _, err := procGetSystemFirmwareTable.Call(
			uintptr(firmwareTableProviderSignature),
			uintptr(FirmwareTableID),
			uintptr(unsafe.Pointer(&pFirmwareTableEnumBuffer[0])),
			uintptr(1),
		)

		if ret == 0 {
			fmt.Println(err)
		}
	}

	fmt.Println(pFirmwareTableEnumBuffer[56:])
}
