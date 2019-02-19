//+build darwin

package networkchangenotifier

import "C"

var userCallback func(data uint64)

//export callback_cgo
func callback_cgo(state C.ulonglong, result C.uint) {
	if userCallback != nil && uint32(result) == 0 {
		userCallback(uint64(state))
	}
}
