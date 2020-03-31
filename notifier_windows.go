//+build windows

package networkchangenotifier

import (
	"os"
	"strings"
)

// Init will register and recv network change events
func (c *NetworkChangeNotifier) Init() error {
	debugChange = strings.Contains(os.Getenv("GODEBUG"), "networkchange=1")
	return ncnInit()
}

// OnNetworkChanged will register user callback function
func (c *NetworkChangeNotifier) OnNetworkChanged(f func(handleUIntPtr uint64)) {
	userCallbackLock.Lock()
	defer userCallbackLock.Unlock()
	userCallback = f
	ncnRegisterCallback()
}

// UnregisterCallback will unregister user callback function
func (c *NetworkChangeNotifier) UnregisterCallback() {
	ncnUnregisterCallback()
	userCallbackLock.Lock()
	defer userCallbackLock.Unlock()
	userCallback = nil
}

// Cleanup will cancel recv network change events
func (c *NetworkChangeNotifier) Cleanup() error {
	return ncnCleanup()
}
