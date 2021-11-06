package sysinfoapi

import (
	"syscall"
	"unsafe"
)

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procEnumSystemFirmwareTables = kernel32.NewProc("EnumSystemFirmwareTables")
)

type FirmwareTable int64

const (
	//ACPI The ACPI firmware table provider.
	ACPI FirmwareTable = 1094930505
	//FIRM The raw firmware table provider. Not supported for UEFI systems; use 'RSMB' instead.
	FIRM FirmwareTable = 1179210317
	//RSMB The raw SMBIOS firmware table provider.
	RSMB FirmwareTable = 1381190978
)

//EnumSystemFirmwareTables
func EnumSystemFirmwareTables(firmwareTableProviderSignature FirmwareTable) (tables []string, err error) {

	ret, _, err := procEnumSystemFirmwareTables.Call(
		uintptr(firmwareTableProviderSignature),
		uintptr(unsafe.Pointer(&make([]uint8, 1)[0])),
		uintptr(1),
	)

	if ret == 0 {
		return
	}

	pFirmwareTableEnumBuffer := make([]uint8, ret)

	ret, _, err = procEnumSystemFirmwareTables.Call(
		uintptr(firmwareTableProviderSignature),
		uintptr(unsafe.Pointer(&pFirmwareTableEnumBuffer[0])),
		uintptr(ret),
	)

	if ret == 0 {
		return
	}

	for i := 0; i < len(pFirmwareTableEnumBuffer); i += 4 {
		tables = append(tables, string(pFirmwareTableEnumBuffer[i:i+4]))
	}

	return
}
