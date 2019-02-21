//+build windows,cgo

package networkchangenotifier

import (
	"sync"

	"golang.org/x/sys/windows"
)

// var userCallback PIPFORWARD_CHANGE_CALLBACK
var userCallback func(handleUIntPtr uint64)
var userCallbackLock sync.RWMutex

func callback_windows(CallerContext windows.Handle, Row windows.Handle, NotificationType int) uintptr {
	userCallbackLock.RLock()
	defer userCallbackLock.RUnlock()
	if userCallback != nil {
		// userCallback(CallerContext, Row, NotificationType)
		userCallback(uint64(Row))
	}
	return uintptr(0)
}
