//+build darwin

package networkchangenotifier

import (
	"os/exec"
	"testing"
	"time"
)

func TestCallbackAddDelIP(t *testing.T) {
	ncn := new(NetworkChangeNotifier)
	err := ncn.Init()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	ncn.OnNetworkChanged(func(dataCFPropertyListRef uint64) {
		t.Log("OnNetworkChanged", dataCFPropertyListRef)
	})

	go func() {
		time.Sleep(1 * time.Second)
		testAddDelIP()
	}()

	time.Sleep(3 * time.Second)
	err = ncn.Cleanup()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// func TestCallbackAddDelRoute(t *testing.T) {
// 	ncn := new(NetworkChangeNotifier)
// 	err := ncn.Init()
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	}
// 	ncn.OnNetworkChanged(func(dataCFPropertyListRef uint64) {
// 		t.Log("OnNetworkChanged", dataCFPropertyListRef)
// 	})
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		testAddDelRoute()
// 	}()
// 	time.Sleep(3 * time.Second)
// 	err = ncn.Cleanup()
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	}
// }

func TestUnReg(t *testing.T) {
	ncn := new(NetworkChangeNotifier)
	err := ncn.Init()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	ncn.OnNetworkChanged(func(dataCFPropertyListRef uint64) {
		t.Log("OnNetworkChanged", dataCFPropertyListRef)
		t.Fail()
	})
	ncn.UnregisterCallback()
	go func() {
		time.Sleep(1 * time.Second)
		testAddDelIP()
	}()
	time.Sleep(3 * time.Second)
	err = ncn.Cleanup()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func testAddDelIP() {
	// if runtime.GOOS == "darwin" {
	exec.Command("/bin/bash", "-c", "source /etc/profile && sudo ifconfig lo0 add 127.0.0.2/8 && sudo ifconfig lo0 delete 127.0.0.2").Run()
	// }
}

// func testAddDelRoute() {
// 	// if runtime.GOOS == "darwin" {
// 	exec.Command("/bin/bash", "-c", "source /etc/profile && sudo route -n add -net 127.0.0.2 -netmask 255.255.255.255 127.0.0.1 && sudo route -n delete -net 127.0.0.2 -netmask 255.255.255.255 127.0.0.1").Run()
// 	// }
// }
