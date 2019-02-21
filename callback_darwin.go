//+build darwin

package networkchangenotifier

import "C"
import "sync"

var userCallback func(dataCFPropertyListRef uint64)
var userCallbackLock sync.RWMutex

//export callback_cgo
func callback_cgo(dataCFPropertyListRef C.ulonglong) {
	userCallbackLock.RLock()
	defer userCallbackLock.RUnlock()
	if userCallback != nil {
		userCallback(uint64(dataCFPropertyListRef))
	}
}
