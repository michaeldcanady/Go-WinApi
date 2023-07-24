package netapi32_test

import (
	"syscall"
	"testing"
	"unsafe"

	mpr "github.com/michaeldcanady/go-winapi/v2/netapi32"
)

func TestNetAddConnection2(t *testing.T) {
	resource := &mpr.NetResource{
		Scope:       mpr.RESOURCE_GLOBALNET,
		Type:        mpr.RESOURCETYPE_DISK,
		DisplayType: mpr.RESOURCEDISPLAYTYPE_GENERIC,
		Usage:       mpr.RESOURCEUSAGE_CONNECTABLE,
		LocalName:   uintptr(0),
		RemoteName:  uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(`\\server\share`))),
		Comment:     uintptr(0),
		Provider:    uintptr(0),
	}

	// Add a valid username, password, and connection flag for your test case
	username := "testuser"
	password := "testpassword"
	flags := mpr.CONNECT_UPDATE_PROFILE

	err := mpr.NetAddConnection2(resource, password, username, flags)
	if err != nil {
		t.Errorf("NetAddConnection2 failed with error: %v", err)
	}
}

func TestNPOpenEnum(t *testing.T) {
	_, err := mpr.NPOpenEnum(mpr.RESOURCE_CONNECTED, mpr.RESOURCETYPE_ANY, mpr.RESOURCEUSAGE_CONNECTABLE, nil)
	if err != nil {
		t.Errorf("NPOpenEnum failed with error: %v", err)
	}
}
