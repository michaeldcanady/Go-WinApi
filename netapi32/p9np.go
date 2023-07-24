//go:build windows
// +build windows

package netapi32

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	p9np, _               = syscall.LoadDLL("p9np.dll")
	procNPOpenEnum, _     = p9np.FindProc("NPOpenEnum")
	procNPEnumResource, _ = p9np.FindProc("NPEnumResourceW")
)

// For more information see https://learn.microsoft.com/en-us/windows/win32/api/npapi/nf-npapi-npopenenum
func NPOpenEnum(scope NetResourceScope, Type NetResourceType, usage NetResourceUsage, netResource *NetResource) (NetResourceHandle, error) {

	var hEnum NetResourceHandle

	ret, _, _ := procNPOpenEnum.Call(
		uintptr(scope),
		uintptr(Type),
		uintptr(usage),
		uintptr(unsafe.Pointer(netResource)),
		uintptr(unsafe.Pointer(&hEnum)),
	)

	if ret != 0 {
		return 0, syscall.Errno(ret)
	}

	return hEnum, nil
}

// For more information see https://learn.microsoft.com/en-us/windows/win32/api/npapi/nf-npapi-npenumresource
func NPEnumResource(hEnum syscall.Handle) ([]NetResource, error) {
	var bufferSize uint32 = 16384
	buffer := make([]byte, bufferSize)
	var entriesRead uint32

	ret, _, _ := procNPEnumResource.Call(
		uintptr(hEnum),
		uintptr(unsafe.Pointer(&entriesRead)),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(bufferSize),
	)

	if ret != 0 {
		return nil, syscall.Errno(ret)
	}

	// Convert the byte slice to a slice of NetResource structs
	var resources []NetResource
	for offset := 0; offset < int(entriesRead); {
		data := buffer[offset:]
		scope := NetResourceScope(binary.LittleEndian.Uint32(data[0:]))
		Type := NetResourceType(binary.LittleEndian.Uint32(data[4:]))
		displayType := NetResourceDisplayType(binary.LittleEndian.Uint32(data[8:]))
		usage := NetResourceUsage(binary.LittleEndian.Uint32(data[12:]))
		localName := string(data[16:])
		remoteName := string(data[24:])
		comment := string(data[32:])
		provider := string(data[40:])

		resource, err := NewNetResource(localName, remoteName, comment, provider, scope, Type, displayType, usage)
		if err != nil {
			return []NetResource{}, fmt.Errorf("unable to marshall net resources: %s", err)
		}

		resources = append(resources, *resource)

		// Each NetResource struct is 44 bytes in size (total size of all fields).
		offset += 44
	}

	return resources, nil
}
