package utils

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNewProcessManager(t *testing.T) {
	var a int = 0
	var pm ProcessManager = NewProcessManager(func(pm ProcessManager, inChan *ProcessChannel, outChan *ProcessChannel) {
		for i := 0; i < 10; i++ {
			a++
		}
	})
	if pm.Status() != DEFAULT {
		t.Fatal("ProcessManager::error : Unable to Intantiate and Init")
	}
	pm.Start()
	time.Sleep(2 * time.Second)
	fmt.Println(pm.Status())
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
	var pm ProcessManager = NewProcessManager(func(pm ProcessManager, inChan *ProcessChannel, outChan *ProcessChannel) {
		for i := 0; i < 10; i++ {
			a++
			*outChan <- fmt.Sprintf("%v", a)
			expectedValue += a
		}
	})
	if pm.Status() != DEFAULT {
		t.Fatal("ProcessManager::error : Unable to Intantiate and Init")
	}
	pm.Start()
	time.Sleep(1 * time.Second)
	var value int = 0
	outChan := pm.GetOutChannel()
	for pm.Running() {
		txt := <-*outChan
		if txt != nil {
			intValue, _ := strconv.Atoi(txt.(string))
			value += intValue
		}

	}
	if pm.Status() != DONE {
		t.Fatal("ProcessManager::error : Unable to Complete tasks")
	}
	if value != expectedValue {
		t.Fatal("ProcessManager::error : Expected <" + strconv.Itoa(expectedValue) + "> Given <" + strconv.Itoa(value) + ">")
	}

}
