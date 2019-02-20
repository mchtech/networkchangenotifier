//+build darwin

package networkchangenotifier

import "C"

var userCallback func(dataCFPropertyListRef uint64)

//export callback_cgo
func callback_cgo(dataCFPropertyListRef C.ulonglong) {
	if userCallback != nil {
		userCallback(uint64(dataCFPropertyListRef))
	}
}
