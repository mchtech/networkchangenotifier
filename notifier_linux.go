//+build linux

package networkchangenotifier

import "C"

// Init will register and recv network change events
func (c *NetworkChangeNotifier) Init() error {
	return ncnInit()
}

// OnNetworkChanged will register user callback function
func (c *NetworkChangeNotifier) OnNetworkChanged(f func(dataCFPropertyListRef uint64)) {
	userCallback = f
	ncnRegisterCallback()
}

// UnregisterCallback will unregister user callback function
func (c *NetworkChangeNotifier) UnregisterCallback() {
	ncnUnregisterCallback()
	userCallback = nil
}

// Cleanup will cancel recv network change events
func (c *NetworkChangeNotifier) Cleanup() error {
	return ncnCleanup()
}
