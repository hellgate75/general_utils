package common

import (
	"fmt"
	"net"
	"testing"
	"time"
)

type fakeAddr struct {
}

func (addr *fakeAddr) Network() string {
	return "tcp"
}
func (addr *fakeAddr) String() string {
	return "192.168.1.22"
}

type fakeNetConn struct {
	remote fakeAddr
}

func (netConn *fakeNetConn) Read(b []byte) (n int, err error) {
	return 0, nil
}
func (netConn *fakeNetConn) Write(b []byte) (n int, err error) {
	return 0, nil
}
func (netConn *fakeNetConn) Close() error {
	return nil
}
func (netConn *fakeNetConn) LocalAddr() net.Addr {
	return &netConn.remote
}
func (netConn *fakeNetConn) RemoteAddr() net.Addr {
	return &netConn.remote
}
func (netConn *fakeNetConn) SetDeadline(t time.Time) error {
	return nil
}
func (netConn *fakeNetConn) SetReadDeadline(t time.Time) error {
	return nil
}
func (netConn *fakeNetConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestNewPortSyndacator(t *testing.T) {
	interval := PortInterval{
		Port(20010),
		Port(20100),
	}
	conn := fakeNetConn{
		fakeAddr{},
	}
	synd := NewPortSyndacator(interval)

	cli, err := synd.Allocate(&conn)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error message : %s", err.Error()))
	}
	if cli == nil {
		t.Fatal("Unexpected Nil Client added to Port Syndacator!!")
	}
	if cli.ServerPort != Port(20010) {
		t.Fatal(fmt.Sprintf("Wrong Port number - Expected <%v> but Given <%v>", Port(20010), cli.ServerPort))
	}
	cliList := synd.List()
	if len(cliList) != 1 {
		t.Fatal(fmt.Sprintf("Wrong List size - Expected <%x> but Given <%x>", 1, len(cliList)))
	}
	cliRef, errP := synd.ClientByPort(cliList[0].ServerPort)
	if errP != nil {
		t.Fatal(fmt.Sprintf("Unexpected error message : %s", errP.Error()))
	}
	if cliRef == nil {
		t.Fatal("Unexpected nil pointer to client reference data!!")
	}
	err = synd.Deallocate(cli)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error message : %s", err.Error()))
	}
	cliList = synd.List()
	if len(cliList) > 0 {
		t.Fatal(fmt.Sprintf("Wrong List size - Expected <%x> but Given <%x>", 0, len(cliList)))
	}
	err = synd.Deallocate(nil)
	if err == nil {
		t.Fatal("Unexpected nil exception for nil pointer to client reference data!!")
	}
}
