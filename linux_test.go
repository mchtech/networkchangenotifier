//+build linux,cgo android,cgo

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

	ncn.OnNetworkChanged(func(dataNLMSG uint64) {
		t.Log("OnNetworkChanged", dataNLMSG)
	})

	go func() {
		time.Sleep(1 * time.Second)
		testAddDelIP()
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

func TestCallbackAddDelRoute(t *testing.T) {
	ncn := new(NetworkChangeNotifier)
	err := ncn.Init()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	ncn.OnNetworkChanged(func(dataNLMSG uint64) {
		t.Log("OnNetworkChanged", dataNLMSG)
	})
	go func() {
		time.Sleep(1 * time.Second)
		testAddDelRoute()
		time.Sleep(1 * time.Second)
		testAddDelRoute()
	}()
	time.Sleep(3 * time.Second)
	err = ncn.Cleanup()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestUnReg(t *testing.T) {
	ncn := new(NetworkChangeNotifier)
	err := ncn.Init()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	ncn.OnNetworkChanged(func(dataNLMSG uint64) {
		t.Log("OnNetworkChanged", dataNLMSG)
		t.Fail()
	})
	ncn.UnregisterCallback()
	go func() {
		time.Sleep(1 * time.Second)
		testAddDelIP()
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
	exec.Command("/bin/bash", "-c", "source /etc/profile && sudo ip addr add 127.0.0.2/8 dev lo && sudo ip addr del 127.0.0.2/8 dev lo").Run()
}

func testAddDelRoute() {
	exec.Command("/bin/bash", "-c", "source /etc/profile && sudo ip route add 127.0.0.2/32 via 127.0.0.1 && sudo ip route del 127.0.0.2/32 via 127.0.0.1").Run()
}
