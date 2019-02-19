//+build darwin

package networkchangenotifier

/*
#include <notify.h>
#include <notify_keys.h>
#include <stdint.h>

extern void callback_cgo(uint64_t state, uint32_t result);

typedef void (*callback_t)(uint64_t, uint32_t);

callback_t callback;

void regCallback(callback_t cb){
   callback = cb;
}

void unregCallback(){
   callback = NULL;
}

int token;

uint registerNetworkChangeEvent()
{
   return notify_register_dispatch(
      kNotifySCNetworkChange,
      &token,
      dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0),
      ^(int t){
        uint64_t state;
        uint32_t result = notify_get_state(t, &state);
        if (callback != NULL){
            callback(result, state);
        }
      }
   );
}

uint unregisterNetworkChangeEvent()
{
   uint32_t ret = notify_cancel(token);
   unregCallback();
   return ret;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ncnInit() error {
	var err error
	ret := uint32(C.registerNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("registerNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}

func ncnRegisterCallback() {
	C.regCallback((C.callback_t)(unsafe.Pointer(C.callback_cgo)))
}

func ncnUnregisterCallback() {
	C.unregCallback()
}

func ncnCleanup() error {
	var err error
	ret := uint32(C.unregisterNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("unregisterNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}
