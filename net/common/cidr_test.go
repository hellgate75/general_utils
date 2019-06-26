package common

import (
	"fmt"
	"testing"
)

func TestGetCIdrInfo(t *testing.T) {
	var cidr string = "192.168.100.14/24"
	info, err := GetCIdrInfo(cidr)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected Exception : %s", err.Error()))
	}
	var expected string = "{192.168.100.14/24 192.168.100.14 192.168.100.255 {192.168.100.0 ffffff00} 8 256 0.0.0.255}"
	var given string = fmt.Sprintf("%v", *info)
	if expected != given {
		t.Fatal(fmt.Sprintf("Wrong value -> Expected: <%s> But Given: <%s>", expected, given))
	}
	cidr = "192.168.100.14"
	_, err = GetCIdrInfo(cidr)
	if err == nil {
		t.Fatal("Unhappy path -> Expect any Exception : Not Arisen")
	}
}
