//go:build windows
// +build windows

package netapi32

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	mpr, _                            = syscall.LoadDLL("mpr.dll")
	procWNetAddConnection, _          = mpr.FindProc("WNetAddConnection")
	procWNetAddConnection2, _         = mpr.FindProc("WNetAddConnection2W")
	procWNetAddConnection3, _         = mpr.FindProc("WNetAddConnection3W")
	procWNetCancelConnection2, _      = mpr.FindProc("WNetCancelConnection2W")
	procWNetCancelConnection, _       = mpr.FindProc("WNetCancelConnectionW")
	procWNetCloseEnum, _              = mpr.FindProc("WNetCloseEnumW")
	procWNetConnectionDialog, _       = mpr.FindProc("WNetConnectionDialogW")
	procWNetConnectionDialog1, _      = mpr.FindProc("WNetConnectionDialog1W")
	procWNetDisconnectDialog1, _      = mpr.FindProc("WNetDisconnectDialog1W")
	procWNetEnumResource, _           = mpr.FindProc("WNetEnumResourceW")
	procWNetGetConnection, _          = mpr.FindProc("WNetGetConnectionW")
	procWNetGetNetworkInformation, _  = mpr.FindProc("WNetGetNetworkInformationW")
	procWNetGetProviderName, _        = mpr.FindProc("WNetGetProviderNameW")
	procWNetGetResourceInformation, _ = mpr.FindProc("WNetGetResourceInformationW")
	procWNetGetResourceParentW, _     = mpr.FindProc("WNetGetResourceParentWW")
	procWNetGetUniversalName, _       = mpr.FindProc("WNetGetUniversalNameW")
	procWNetGetUser, _                = mpr.FindProc("WNetGetUserW")
	procWNetOpenEnum, _               = mpr.FindProc("WNetOpenEnumW")
	procWNetRestoreConnection, _      = mpr.FindProc("WNetRestoreConnectionW")
	procWNetUseConnection, _          = mpr.FindProc("WNetUseConnectionW")
)

// NetAddConnection2 creates a network connection using the provided network resource
// information and connection options.
//
// The function attempts to establish a network connection to a remote resource,
// using the given credentials for authentication (if required). The `resource`
// parameter specifies the network resource to connect to, and the `pass` and `user`
// parameters contain the password and username for authentication.
//
// The `flags` parameter is a set of connection options that control the behavior of
// the network connection. These options can be combined using bitwise OR operations.
//
// Note that the password and username provided in `pass` and `user` parameters
// will be converted to UTF-16 format, as expected by the underlying Windows API.
//
// The function returns an error if the network connection cannot be established,
// or if there is an issue converting the password or username to UTF-16 format.
// More detailed information at: https://docs.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetaddconnection2w
func NetAddConnection2(resource *NetResource, pass, user string, flags ConnectionFlag) error {

	var (
		_pass *uint16
		_user *uint16
		err   error
	)

	if pass != "" {
		_pass, err = syscall.UTF16PtrFromString(pass)
		if err != nil {
			return fmt.Errorf("unable to convert pass to UTF16Ptr: %s", err)
		}
	} else {
		_pass = nil
	}

	if user != "" {
		_user, err = syscall.UTF16PtrFromString(user)
		if err != nil {
			return fmt.Errorf("unable to convert user to UTF16Ptr: %s", err)
		}
	} else {
		_user = nil
	}

	_password := uintptr(unsafe.Pointer(_pass))
	_username := uintptr(unsafe.Pointer(_user))
	_resource := uintptr(unsafe.Pointer(resource))
	connectionFlags := uintptr(flags)

	ret, _, _ := procWNetAddConnection2.Call(_resource, _password, _username, connectionFlags)

	if ret != 0 {
		return syscall.Errno(ret)
	}
	return nil
}

// NetAddConnection3 function makes a connection to a network resource. The function can redirect a local device to the network resource.
//
// The NetAddConnection3 function is similar to the WNetAddConnection2 function.
// The main difference is that NetAddConnection3 has an additional parameter,
// a handle to a window that the provider of network resources can use as an owner window for dialog boxes.
// The WNetAddConnection2 function and the NetAddConnection3 function supersede the WNetAddConnection function.
func NetAddConnection3(handle NetResourceHandle, resource *NetResource, pass, user string, flags ConnectionFlag) error {

	_pass, err := syscall.UTF16PtrFromString(pass)
	if err != nil {
		return fmt.Errorf("unable to convert password to UTF16Ptr: %s", err)
	}
	_user, err := syscall.UTF16PtrFromString(user)
	if err != nil {
		return fmt.Errorf("unable to convert username to UTF16Ptr: %s", err)
	}

	password := uintptr(unsafe.Pointer(_pass))
	username := uintptr(unsafe.Pointer(_user))
	connectionFlags := uintptr(flags)
	_handle := uintptr(handle)

	ret, _, _ := procWNetAddConnection3.Call(_handle, uintptr(unsafe.Pointer(&resource)), password, username, connectionFlags)

	if ret != 0 {
		return syscall.Errno(ret)
	}
	return nil
}

func WNetAddConnection() error {
	return errors.New("WNetAddConnection is not implemented")
}
func WNetCancelConnection() error {
	return errors.New("WNetCancelConnection is not implemented")
}
func WNetCloseEnum() error {
	return errors.New("WNetCloseEnum is not implemented")
}
func WNetConnectionDialog() error {
	return errors.New("WNetConnectionDialog is not implemented")
}
func WNetConnectionDialog1() error {
	return errors.New("WNetConnectionDialog1 is not implemented")
}
func WNetDisconnectDialog1() error {
	return errors.New("WNetDisconnectDialog1 is not implemented")
}
func WNetEnumResource() error {
	return errors.New("WNetEnumResource is not implemented")
}
func WNetGetConnection() error {
	return errors.New("WNetGetConnection is not implemented")
}
func WNetGetNetworkInformation() error {
	return errors.New("WNetGetNetworkInformation is not implemented")
}
func WNetGetProviderName() error {
	return errors.New("WNetGetProviderName is not implemented")
}
func WNetGetResourceInformation() error {
	return errors.New("WNetGetResourceInformation is not implemented")
}
func WNetGetResourceParentW() error {
	return errors.New("WNetGetResourceParentW is not implemented")
}
func WNetGetUniversalName() error {
	return errors.New("WNetGetUniversalName is not implemented")
}
func WNetGetUser() error {
	return errors.New("WNetGetUser is not implemented")
}
func WNetOpenEnum() error {
	return errors.New("WNetOpenEnum is not implemented")
}
func WNetRestoreConnection() error {
	return errors.New("WNetRestoreConnection is not implemented")
}
func WNetUseConnection() error {
	return errors.New("WNetUseConnection is not implemented")
}
