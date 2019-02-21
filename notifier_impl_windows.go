//+build windows,cgo

package networkchangenotifier

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var iphlpapi = windows.NewLazySystemDLL("iphlpapi.dll")
var notifyRouteChange2 = iphlpapi.NewProc("NotifyRouteChange2")
var cancelMibChangeNotify2 = iphlpapi.NewProc("CancelMibChangeNotify2")

var h windows.Handle

func ncnInit() error {
	var err error
	ret, err := NotifyRouteChange2(
		uint16(windows.AF_UNSPEC),
		callback_windows,
		0,
		0,
		&h,
	)
	if ret == windows.NO_ERROR {
		err = nil
	}
	return err
}

func ncnRegisterCallback() {
}

func ncnUnregisterCallback() {
}

func ncnCleanup() error {
	var err error
	ret, err := CancelMibChangeNotify2(
		h,
	)
	if ret == windows.NO_ERROR {
		err = nil
	}
	return err
}

type PIPFORWARD_CHANGE_CALLBACK func(CallerContext windows.Handle, Row windows.Handle, NotificationType int) uintptr

func NotifyRouteChange2(
	AddressFamily uint16,
	Callback PIPFORWARD_CHANGE_CALLBACK,
	CallerContext windows.Handle,
	InitialNotification uint8,
	NotificationHandle *windows.Handle,
) (uint32, error) {
	cb := syscall.NewCallback(Callback)
	ret, _, err := notifyRouteChange2.Call(
		uintptr(AddressFamily),
		cb,
		uintptr(unsafe.Pointer(&CallerContext)),
		uintptr(InitialNotification),
		uintptr(unsafe.Pointer(NotificationHandle)),
	)
	return uint32(ret), err
}

func CancelMibChangeNotify2(NotificationHandle windows.Handle) (uint32, error) {
	ret, _, err := cancelMibChangeNotify2.Call(
		uintptr(NotificationHandle),
	)
	return uint32(ret), err
}
