//+build linux

package networkchangenotifier

import "C"
import "sync"

var userCallback func(dataNLMSG uint64)
var userCallbackLock sync.RWMutex

//export callback_cgo
func callback_cgo(dataNLMSG C.ulonglong) {
	userCallbackLock.RLock()
	defer userCallbackLock.RUnlock()
	if userCallback != nil {
		userCallback(uint64(dataNLMSG))
	}
}
