//+build windows

package networkchangenotifier

import "golang.org/x/sys/windows"

// var userCallback PIPFORWARD_CHANGE_CALLBACK
var userCallback func(handleUIntPtr uint64)

func callback_windows(CallerContext windows.Handle, Row windows.Handle, NotificationType int) uintptr {
	if userCallback != nil {
		// userCallback(CallerContext, Row, NotificationType)
		userCallback(uint64(Row))
	}
	return uintptr(0)
}
