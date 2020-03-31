package networkchangenotifier

// INetworkChangeNotifier is
type INetworkChangeNotifier interface {
	Init() error
	OnNetworkChanged(func(Ptr uint64))
	UnregisterCallback()
	Cleanup() error
}

// NetworkChangeNotifier is
type NetworkChangeNotifier struct {
}

var debugChange bool
