//go:build windows
// +build windows

package netapi32

import "syscall"

type (
	ConnectionFlag         uint32
	NetResourceScope       uint32
	NetResourceType        uint32
	NetResourceDisplayType uint32
	NetResourceUsage       uint32
	NetResourceHandle      syscall.Handle
)

const (
	CONNECT_UPDATE_PROFILE ConnectionFlag = 0x00000001
	CONNECT_UPDATE_RECENT  ConnectionFlag = 0x00000002
	CONNECT_TEMPORARY      ConnectionFlag = 0x00000004
	CONNECT_INTERACTIVE    ConnectionFlag = 0x00000008
	CONNECT_PROMPT         ConnectionFlag = 0x00000010
	CONNECT_REDIRECT       ConnectionFlag = 0x00000080
	CONNECT_CURRENT_MEDIA  ConnectionFlag = 0x00000200
	CONNECT_COMMANDLINE    ConnectionFlag = 0x00000800
	CONNECT_CMD_SAVECRED   ConnectionFlag = 0x00001000
	CONNECT_CRED_RESET     ConnectionFlag = 0x00002000

	//RESOURCE_CONNECTED Current connections to network resources.
	RESOURCE_CONNECTED NetResourceScope = iota
	//RESOURCE_GLOBALNET All network resources. These may or may not be connected.
	RESOURCE_GLOBALNET
	//RESOURCE_CONTEXT The network resources associated with the user's current and default network context. The meaning of this is provider-specific.
	RESOURCE_CONTEXT

	//RESOURCETYPE_DISK The resource is a shared disk volume.
	RESOURCETYPE_DISK NetResourceType = iota
	//RESOURCETYPE_PRINT The resource is a shared printer.
	RESOURCETYPE_PRINT
	//RESOURCETYPE_ANY The resource matches more than one type, for example, a container of both print and disk resources, or a resource which is neither print or disk.
	RESOURCETYPE_ANY

	//RESOURCEDISPLAYTYPE_NETWORK The resource is a network provider.
	RESOURCEDISPLAYTYPE_NETWORK NetResourceDisplayType = iota
	//RESOURCEDISPLAYTYPE_DOMAIN The resource is a collection of servers.
	RESOURCEDISPLAYTYPE_DOMAIN
	//RESOURCEDISPLAYTYPE_SERVER The resource is a server.
	RESOURCEDISPLAYTYPE_SERVER
	//RESOURCEDISPLAYTYPE_SHARE The resource is a share point.
	RESOURCEDISPLAYTYPE_SHARE
	//RESOURCEDISPLAYTYPE_DIRECTORY The resource is a directory.
	RESOURCEDISPLAYTYPE_DIRECTORY
	//RESOURCEDISPLAYTYPE_GENERIC The resource type is unspecified.
	//This value is used by network providers that do not specify resource types.
	RESOURCEDISPLAYTYPE_GENERIC

	//RESOURCEUSAGE_CONNECTABLE You can connect to the resource by calling NPAddConnection.
	//If dwType is RESOURCETYPE_DISK, then, after you have connected to the resource, you can use the file system APIs, such as FindFirstFile, and FindNextFile, to enumerate any files and directories the resource contains.
	RESOURCEUSAGE_CONNECTABLE NetResourceUsage = iota
	//RESOURCEUSAGE_CONTAINER The resource is a container for other resources that can be enumerated by means of the NPOpenEnum, NPEnumResource, and NPCloseEnum functions.
	//
	//The container may, however, be empty at the time the enumeration is made. In other words, the first call to NPEnumResource may return WN_NO_MORE_ENTRIES.
	RESOURCEUSAGE_CONTAINER
)
