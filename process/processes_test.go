package process

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNewProcessManager(t *testing.T) {
	var a int = 0
	var pm ProcessManager = NewProcessManager(func(pm ProcessManager, inChan *ProcessChannel, outChan *ProcessChannel) error {
		for i := 0; i < 10; i++ {
			a++
		}
		return nil
	})
	if pm.Status() != DEFAULT {
		t.Fatal("ProcessManager::error : Unable to Intantiate and Init")
	}
	pm.Start()
	for pm.Running() {
		time.Sleep(500 * time.Millisecond)
	}
	//	fmt.Println(fmt.Sprintf("1.State: %v", pm.Status()))
	if pm.Status() != DONE {
		t.Fatal("ProcessManager::error : Unable to Complete tasks")
	}
	if a != 10 {
		t.Fatal("ProcessManager::error : Expected <10> Given <" + strconv.Itoa(a) + ">")
	}

}

func TestProcessManagerChannels(t *testing.T) {
	var a int = 0
	var expectedValue int = 0
	var pm ProcessManager = NewProcessManager(func(pm ProcessManager, inChan *ProcessChannel, outChan *ProcessChannel) error {
		for i := 0; i < 10; i++ {
			a++
			*outChan <- fmt.Sprintf("%v", a)
			expectedValue += a
		}
		return errors.New("Test Error!!")
	})
	if pm.Status() != DEFAULT {
		t.Fatal("ProcessManager::error : Unable to Intantiate and Init")
	}
	pm.Start()
	var value int = 0
	outChan := pm.GetOutChannel()
	for pm.Running() {
		txt := <-*outChan
		if txt != nil {
			intValue, _ := strconv.Atoi(txt.(string))
			value += intValue
		}

	}
	//	fmt.Println(fmt.Sprintf("2.State: %v", pm.Status()))
	if pm.Status() != ERROR {
		t.Fatal("ProcessManager::error : Unable to Complete tasks")
	}
	if value != expectedValue {
		t.Fatal("ProcessManager::error : Expected <" + strconv.Itoa(expectedValue) + "> Given <" + strconv.Itoa(value) + ">")
	}

}
