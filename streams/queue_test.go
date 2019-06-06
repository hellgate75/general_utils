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

func TestMessageQueueEchoBrisge(t *testing.T) {
	var message common.Message = "Test Queue"
	var outcome common.Message = ""
	var readTimeout time.Duration = 1 * time.Second
	var echoReadTimeout time.Duration = 1 * time.Second
	messageQueue := NewMessageQueue(readTimeout)

	var echoBridgeChan chan common.Message = make(chan common.Message)
	var echoMessageBridge MessageBridge = NewEchoMessageBridge(&echoBridgeChan, readTimeout, echoReadTimeout)

	defer func() {
		messageQueue.Stop()
		echoMessageBridge.Close()
	}()

	messageQueue.AddMessageBridge(&echoMessageBridge)
	messageQueue.Start()
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
