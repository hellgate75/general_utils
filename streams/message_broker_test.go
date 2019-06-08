package streams

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
	"time"
)

func TestNewNoTransitionMessageBrokerWithTimeOut(t *testing.T) {
	var message common.Message = "Test NewMessageBroker"
	var outcome common.Message = ""
	var err error = nil
	queue := NewNoTransitionMessageBroker(time.Second)
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

func TestNewNoTransitionMessageBrokerWithoutTimeOut(t *testing.T) {
	var message common.Message = "Test NewMessageBroker"
	var outcome common.Message = ""
	var err error = nil
	queue := NewNoTransitionMessageBroker(0 * time.Second)
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

func TestNewMessageBrokerEchoBridge(t *testing.T) {
	var message common.Message = "Test NewMessageBroker"
	var outcome common.Message = ""
	var readTimeout time.Duration = 1 * time.Second
	var echoReadTimeout time.Duration = 1 * time.Second
	messageBroker := NewMessageBroker(readTimeout)

	var echoBridgeChan chan common.Message = make(chan common.Message)
	var echoMessageBridge MessageBridge = NewEchoMessageBridge(&echoBridgeChan, readTimeout, echoReadTimeout)

	defer func() {
		messageBroker.Stop()
		echoMessageBridge.Close()
	}()

	messageBroker.AddMessageBridge(&echoMessageBridge)
	messageBroker.Start()
	echoMessageBridge.Open()
	echoBridgeChan <- message
	select {
	case msg := <-echoBridgeChan:
		outcome = msg
	}
	if outcome == "" {
		t.Fatal("Unable to receive any suitable message")
	} else if outcome != message {
		t.Fatal("Unable to receive any suitable message")
	}
}
