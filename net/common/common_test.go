package common

import (
	"fmt"
	"testing"
)

func TestServerTypes(t *testing.T) {
	if TCP != ServerType(1) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerType(1), TCP))
	}
	if UDP != ServerType(2) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerType(2), UDP))
	}
	if REST != ServerType(3) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerType(3), REST))
	}
	if CONTENT != ServerType(4) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerType(4), CONTENT))
	}
}

func TestServerStateSignals(t *testing.T) {
	if RUNNING != ServerStateSignal(500) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(500), RUNNING))
	}
	if PAUSED != ServerStateSignal(501) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(501), PAUSED))
	}
	if DISCONNECTING != ServerStateSignal(502) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(502), DISCONNECTING))
	}
	if REFUSED != ServerStateSignal(503) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(503), REFUSED))
	}
	if FAILURE != ServerStateSignal(504) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(504), FAILURE))
	}
	if STOPPING != ServerStateSignal(505) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(505), STOPPING))
	}
	if STOPPED != ServerStateSignal(506) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", ServerStateSignal(506), STOPPED))
	}
}
