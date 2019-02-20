//+build windows

package networkchangenotifier

// Init will register and recv network change events
func (c *NetworkChangeNotifier) Init() error {
	return ncnInit()
}

// OnNetworkChanged will register user callback function
func (c *NetworkChangeNotifier) OnNetworkChanged(f func(handleUIntPtr uint64)) {
	ncnRegisterCallback()
	userCallback = f
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
