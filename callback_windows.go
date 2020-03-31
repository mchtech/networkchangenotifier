//+build windows

package networkchangenotifier

import (
	"log"
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
		userCallback(uint64(Row))
		if debugChange {
			log.Println("networkchange:", CallerContext, Row, NotificationType)
		}
	}
	return uintptr(0)
}
