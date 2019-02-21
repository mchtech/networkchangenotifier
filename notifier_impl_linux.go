//+build !windows,!darwin

package networkchangenotifier

import (
	"sync"
	"time"

	"github.com/vishvananda/netlink"
)

type nullstruct struct {
}

func ncnInit() error {
	var err error
	stopped = false
	done = make(chan struct{})
	return err
}

var done chan struct{}
var once sync.Once
var stopped bool

func ncnRegisterCallback() {
	once.Do(func() {
		go func() {
			ch := make(chan netlink.RouteUpdate)
			if err := netlink.RouteSubscribeWithOptions(ch, done, netlink.RouteSubscribeOptions{
				ErrorCallback: func(err error) {
					if err != nil {
						stopped = true
						return
					}
				},
			}); err != nil {
				return
			}
			for {
				timeout := time.After(time.Second)
				select {
				case update := <-ch:
					callback_linux(&update)
				case <-timeout:
					if stopped {
						break
					}
					continue
				}
			}
		}()
	})
}

func ncnUnregisterCallback() {
}

func ncnCleanup() error {
	var err error
	// ncnUnregisterCallback()
	// done <- nullstruct{}
	// close(done)
	return err
}
