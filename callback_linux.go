//+build !windows,!darwin

package networkchangenotifier

import (
	"encoding/json"
	"fmt"
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
		if bt, err := json.Marshal(dataNLMSG); err == nil {
			fmt.Println(string(bt))
		} else {
			fmt.Println(err)
		}
	}
}
