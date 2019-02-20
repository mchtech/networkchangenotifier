//+build darwin

package networkchangenotifier

/*
#cgo darwin LDFLAGS: -framework SystemConfiguration -framework CoreFoundation
#include <stdint.h>
#include <SystemConfiguration/SCDynamicStore.h>
#include <SystemConfiguration/SystemConfiguration.h>

extern void callback_cgo(uint64_t dataCFPropertyListRef);

typedef void (*callback_t)(uint64_t);

callback_t callback;

SCDynamicStoreRef store;

void regCallback(callback_t cb){
   callback = cb;

   CFRunLoopSourceRef rlSrc = SCDynamicStoreCreateRunLoopSource(kCFAllocatorDefault, store, 0);
   CFRunLoopAddSource(CFRunLoopGetCurrent(), rlSrc, kCFRunLoopDefaultMode);
   CFRelease(rlSrc);

   CFRunLoopRun();
}

void unregCallback(){
   CFRunLoopStop(CFRunLoopGetCurrent());
   callback = NULL;
}

void DynamicStoreCallBack(SCDynamicStoreRef store, CFArrayRef changedKeys, void *info)
{
   CFPropertyListRef data =  SCDynamicStoreCopyValue(store, (CFStringRef)CFArrayGetValueAtIndex(changedKeys, 0));
   // CFShow(data);
   if (callback != NULL){
      callback((uint64_t)data);
   }
   CFRelease(data);
}

int registerNetworkChangeEvent()
{
   store = SCDynamicStoreCreate(
      kCFAllocatorDefault,
      CFBundleGetIdentifier(CFBundleGetMainBundle()),
      DynamicStoreCallBack,
      NULL
   );
   CFStringRef strs[1] = {
      CFSTR("State:/Network/Interface/.+/IPv.+"),
   };
   CFArrayRef confArray = CFArrayCreate(kCFAllocatorDefault, (const void**)strs, 1, &kCFTypeArrayCallBacks);
   // CFShow(confArray);
   if(!SCDynamicStoreSetNotificationKeys(store, NULL, confArray))
   {
      CFRelease(confArray);
      CFRelease(store);
      return SCError();
   }
   CFRelease(confArray);
   return 0;
}

int unregisterNetworkChangeEvent()
{
   unregCallback();
   CFRelease(store);
   store = NULL;
   return 0;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ncnInit() error {
	var err error
	ret := int32(C.registerNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("registerNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}

func ncnRegisterCallback() {
	go C.regCallback((C.callback_t)(unsafe.Pointer(C.callback_cgo)))
}

func ncnUnregisterCallback() {
	C.unregCallback()
}

func ncnCleanup() error {
	var err error
	ret := int32(C.unregisterNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("unregisterNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}
