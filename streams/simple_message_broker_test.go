package streams

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
	"time"
)

func TestNewSimpleMessageBroker(t *testing.T) {
	var message common.Type = "TestPipeline"
	messageBroker := NewSimpleMessageBroker()
	if !messageBroker.IsOpen() {
		messageBroker.Open()
	}
	defer func(messageBroker SimpleMessageBroker) {
		messageBroker.Close()
	}(messageBroker)
	go func(messageBroker SimpleMessageBroker) {
		messageBroker.Write(message)
	}(messageBroker)
	msg, err := messageBroker.Read(time.Second)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to Read from the Pipeline : err %s", err.Error()))
	} else if msg != message {
		t.Fatal(fmt.Sprintf("Wrong read message from the Pipeline Expected <%v> Received <%v>", message, msg))
	}
}
