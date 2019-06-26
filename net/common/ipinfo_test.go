package common

import (
	"fmt"
	"testing"
)

func TestGetIpAddressInfo(t *testing.T) {
	var ip_address string = "192.168.100.14"
	info, err := GetIpAddressInfo(ip_address)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected Exception : %s", err.Error()))
	}
	var expected string = "{192.168.100.14 192.168.100.14 255.255.255.0 false true false false false false}"
	var given string = fmt.Sprintf("%v", *info)
	if expected != given {
		t.Fatal(fmt.Sprintf("Wrong value -> Expected: <%s> But Given: <%s>", expected, given))
	}
	ip_address = "192.168.100.14/24"
	_, err = GetIpAddressInfo(ip_address)
	if err == nil {
		t.Fatal("Unhappy path -> Expect any Exception : Not Arisen")
	}
}

func TestGetIpAddressInfoUnhappyPath(t *testing.T) {
	var ip_address string = "dsfdsdfsdf"
	_, err := GetIpAddressInfo(ip_address)
	if err == nil {
		t.Fatal("Unexpected nil exception !!!")
	}
}
