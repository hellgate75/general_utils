package streams

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
	"time"
)

func TestNewQueueWithTimeOut(t *testing.T) {
	var message common.Message = "Test Queue"
	var outcome common.Message = ""
	var err error = nil
	queue := NewQueue(time.Second)
	go func() {
		time.Sleep(350 * time.Millisecond)
		queue.Write(message)
	}()
	outcome, err = queue.Read()
	if err != nil {
		t.Fatal(err.Error())
	} else if outcome != message {
		t.Fatal(fmt.Sprintf("Wrong received object expected <%v> received <%v>", message, outcome))
	}
}

func TestNewQueueWithoutTimeOut(t *testing.T) {
	var message common.Message = "Test Queue"
	var outcome common.Message = ""
	var err error = nil
	queue := NewQueue(0 * time.Second)
	go func() {
		time.Sleep(350 * time.Millisecond)
		queue.Write(message)
	}()
	outcome, err = queue.Read()
	if err != nil {
		t.Fatal(err.Error())
	} else if outcome != message {
		t.Fatal(fmt.Sprintf("Wrong received object expected <%v> received <%v>", message, outcome))
	}
}
