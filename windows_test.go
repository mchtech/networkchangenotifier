//+build windows

package networkchangenotifier

import (
	"fmt"
	"net"
	"os"
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
	ncn.OnNetworkChanged(func(HandleUIntPtr uint64) {
		t.Log("OnNetworkChanged", Row)
	})
	go testAddDelIP()
	time.Sleep(2 * time.Second)
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
	ncn.OnNetworkChanged(func(HandleUIntPtr uint64) {
		t.Log("OnNetworkChanged", Row)
	})
	go testAddDelRoute()
	time.Sleep(2 * time.Second)
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
	ncn.OnNetworkChanged(func(HandleUIntPtr uint64) {
		t.Log("OnNetworkChanged", Row)
		t.Fail()
	})
	ncn.UnregisterCallback()
	go testAddDelIP()
	time.Sleep(2 * time.Second)
	err = ncn.Cleanup()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// run as Administrator Please
func testAddDelIP() {
	cmd1 := exec.Command(`C:\Windows\System32\netsh.exe`, "interface", "ipv4", "add", "address", "1", "127.0.0.2", "255.0.0.0")
	cmd1.Stderr = os.Stderr
	cmd1.Stdout = os.Stdout
	fmt.Println(cmd1.Run())

	cmd2 := exec.Command(`C:\Windows\System32\netsh.exe`, "interface", "ipv4", "delete", "address", "1", "127.0.0.2")
	cmd2.Stderr = os.Stderr
	cmd2.Stdout = os.Stdout
	fmt.Println(cmd2.Run())
}

// run as Administrator Please
func testAddDelRoute() {
	var cip string
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		ip, _, _ := net.ParseCIDR(addr.String())
		if ip != nil && ip.IsGlobalUnicast() && ip.To4() != nil {
			cip = ip.String()
			break
		}
	}
	cmd1 := exec.Command(`C:\Windows\System32\route.exe`, "add", "127.0.0.2", "mask", "255.255.255.255", cip)
	cmd1.Stderr = os.Stderr
	cmd1.Stdout = os.Stdout
	fmt.Println(cmd1.Run())

	cmd2 := exec.Command(`C:\Windows\System32\route.exe`, "delete", "127.0.0.2", "mask", "255.255.255.255")
	cmd2.Stderr = os.Stderr
	cmd2.Stdout = os.Stdout
	fmt.Println(cmd2.Run())
}
