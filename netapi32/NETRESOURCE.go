//go:build windows
// +build windows

package netapi32

import (
	"fmt"
	"syscall"
	"unsafe"
)

type NetResource struct {
	dwScope       NetResourceScope
	dwType        NetResourceType
	dwDisplayType NetResourceDisplayType
	dwUsage       NetResourceUsage
	lpLocalName   uintptr
	lpRemoteName  uintptr
	lpComment     uintptr
	lpProvider    uintptr
}

// NewNetResource creates a new NetResource struct initialized with the provided values.
//
// Parameters:
//
//	localName: The local name for the network resource (optional). Set to an empty string or nil for no local name.
//	remoteName: The remote name for the network resource (required). Should be a valid UNC path (e.g., \\server\share).
//	comment: A comment associated with the network resource (optional). Set to an empty string or nil for no comment.
//	provider: The network provider that owns the resource (optional). Set to an empty string or nil for no provider.
//	scope: The scope of the network resource (e.g., RESOURCE_CONNECTED, RESOURCE_GLOBALNET, RESOURCE_CONTEXT).
//	Type: The type of the network resource (e.g., RESOURCETYPE_DISK, RESOURCETYPE_PRINT, RESOURCETYPE_ANY).
//	displayType: The display type of the network resource (e.g., RESOURCEDISPLAYTYPE_NETWORK, RESOURCEDISPLAYTYPE_SHARE, RESOURCEDISPLAYTYPE_GENERIC).
//	usage: The usage of the network resource (e.g., RESOURCEUSAGE_CONNECTABLE, RESOURCEUSAGE_CONTAINER).
//
// Returns:
//
//	A new NetResource struct initialized with the provided values.
func NewNetResource(localName, remoteName, comment, provider string, scope NetResourceScope, Type NetResourceType, displayType NetResourceDisplayType, usage NetResourceUsage) (*NetResource, error) {

	_localName, err := syscall.UTF16PtrFromString(localName)
	if err != nil {
		return nil, fmt.Errorf("unable to convert localName to UTF16Ptr: %s", err)
	}

	_remoteName, err := syscall.UTF16PtrFromString(remoteName)
	if err != nil {
		return nil, fmt.Errorf("unable to convert remoteName to UTF16Ptr: %s", err)
	}

	_comment, err := syscall.UTF16PtrFromString(comment)
	if err != nil {
		return nil, fmt.Errorf("unable to convert comment to UTF16Ptr: %s", err)
	}

	_provider, err := syscall.UTF16PtrFromString(provider)
	if err != nil {
		return nil, fmt.Errorf("unable to convert provider to UTF16Ptr: %s", err)
	}

	return &NetResource{
		//dwScope:       scope, NEED MORE WORK
		//dwType:        Type, NEED MORE WORK
		//dwDisplayType: displayType, NEED MORE WORK
		//dwUsage:       usage, NEED MORE WORK
		lpLocalName:  uintptr(unsafe.Pointer(_localName)),
		lpRemoteName: uintptr(unsafe.Pointer(_remoteName)),
		lpComment:    uintptr(unsafe.Pointer(_comment)),
		lpProvider:   uintptr(unsafe.Pointer(_provider)),
	}, nil
}
