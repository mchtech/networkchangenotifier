//+build linux

package networkchangenotifier

import "C"

var userCallback func(dataNLMSG uint64)

//export callback_cgo
func callback_cgo(dataNLMSG C.ulonglong) {
	if userCallback != nil {
		userCallback(uint64(dataNLMSG))
	}
}
