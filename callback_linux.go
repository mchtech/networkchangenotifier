//+build !windows,!darwin

package networkchangenotifier

import (
	"sync"
	"unsafe"

	"github.com/vishvananda/netlink"
)

var userCallback func(dataNLMSG uint64)
var userCallbackLock sync.RWMutex

func callback_linux(dataNLMSG *netlink.RouteUpdate) {
	userCallbackLock.RLock()
	defer userCallbackLock.RUnlock()
	if userCallback != nil {
		userCallback(uint64(uintptr(unsafe.Pointer(dataNLMSG))))
	}
}
