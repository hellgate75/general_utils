package streams

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
	"time"
)

func TestNewPipeline(t *testing.T) {
	var message common.Type = "TestPipeline"
	pipeline := NewPipeline()
	if !pipeline.IsOpen() {
		pipeline.Open()
	}
	defer func(pipeline Pipeline) {
		pipeline.Close()
	}(pipeline)
	go func(pipeline Pipeline) {
		pipeline.Write(message)
	}(pipeline)
	msg, err := pipeline.Read(time.Second)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to Read from the Pipeline : err %s", err.Error()))
	} else if msg != message {
		t.Fatal(fmt.Sprintf("Wrong read message from the Pipeline Expected <%v> Received <%v>", message, msg))
	}
}
