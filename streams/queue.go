package streams

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"runtime"
	"sync"
	"time"
)

type BridgeType int

const (
	ReadOnly  BridgeType = 1021
	WriteOnly BridgeType = 1022
	ReadWrite BridgeType = 1023
)

type MessageBridge interface {
	// Retrieve the Bridge Name.
	//
	// Returns:
	//   string Name defining the Bridge kind, features or brand
	Name() string
	// Opens the Bridge.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Open() error
	// Closes the Bridge.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Close() error
	// Reveal is the Bridge conenction is open.
	//
	// Returns:
	//   bool Open state of the Bridge connection
	IsOpen() bool
	// Returns the Exchange Channel.
	//
	// Returns:
	// *chan common.Message channel to communicate in/out the bridge
	GetInOutChannel() *chan common.Message
	// Tells the queue which purpose has the message bridge.
	//
	// Returns:
	//  streams.BridgeType describe nature and behaviour of the message bridge
	GetBridgeType() BridgeType
}

// Interface describes Queue features
type Queue interface {
	//Initialize the message queue flow
	init()
	// Writes Message into the message queue.
	//
	// Parameters:
	//   object (common.Type) The message to write into the queue
	// Returns:
	//   error Any suitable error risen during code execution
	Write(common.Type) error
	// Reads Message from the message queue.
	//
	// Returns:
	//  ( common.Type object read from the queue,
	//   error Any suitable error risen during code execution )
	Read() (common.Type, error)
}

type queueStruct struct {
	sync.RWMutex
	inOutChan   chan common.Type
	initialized bool
	TimeOut     time.Duration
}

func (q *queueStruct) init() {
	if q.initialized {
		return
	}
	var IOchan chan common.Type = make(chan common.Type)
	q.inOutChan = IOchan
	q.initialized = true
}

func (q *queueStruct) Write(m common.Type) error {
	q.inOutChan <- m
	return nil
}

func (q *queueStruct) Read() (common.Type, error) {
	var val common.Type
	if q.TimeOut == 0 {
		select {
		case msg := <-q.inOutChan:
			val = msg
		}
	} else {
		//		q.RLock()
		select {
		case msg := <-q.inOutChan:
			val = msg
		case <-time.After(q.TimeOut):
			return nil, errors.New("MessageQueue::Read::err Timeout reached")
		}
		//		q.RUnlock()
	}
	return val, nil
}

//Creates a New Queue
//
// Parameters:
//   messageReadTimeout (time.Duration) Timeout before reset the listening (infinite if <= 0)
// Returns:
//   Queue New brand queue
func NewQueue(messageReadTimeout time.Duration) Queue {
	var q queueStruct = queueStruct{
		TimeOut: messageReadTimeout,
	}
	runtime.SetFinalizer(&q, func(qs *queueStruct) {
		defer func() {
			err := recover()
			if logger != nil {
				logger.Error(err.(error))
			}
		}()
		close(qs.inOutChan)
	})
	q.init()
	return &q

}

// Interface describes Message Queue features
type MessageQueue interface {
	//Initialize the message queue flow
	init()
	// Add a Message Bridge into the message queue. The queue will broadcast messages
	// to any of the bridges, in read, write or read/write mode, accordingly to bridge nature
	//
	// Parameters:
	//   bridge (*streams.MessageBridge) The bridge pointer to be linked to the message queue broadcast chain
	// Returns:
	//   error Any suitable error risen during code execution
	AddMessageBridge(*MessageBridge) error
	// Retrieve the link of the linked message bridges
	//
	// Returns:
	//   []*streams.MessageBridge List of linked message bridges
	ListMessageBridges() []*MessageBridge
	// Start the message queue.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Start() error
	// Stop the message queue.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Stop() error
	// Retrive running state of the message queue.
	//
	// Returns:
	//   bool Running state of the message queue
	IsRunning() bool
}

type messageQueueStruct struct {
	sync.RWMutex
	internalBridgeChannel chan common.Message
	bridges               []*MessageBridge
	initialized           bool
	started               bool
	TimeOut               time.Duration
	internalTimeOut       time.Duration
}

func (q *messageQueueStruct) init() {
	if q.initialized {
		return
	}
	if q.TimeOut < 0 {
		q.TimeOut = 0
	}
	q.internalBridgeChannel = make(chan common.Message)
	q.internalTimeOut = 10 * time.Second
	q.initialized = true
}

func (q *messageQueueStruct) AddMessageBridge(bridge *MessageBridge) error {
	if bridge == nil {
		return errors.New("MessageQueue::AddMessageBridge::err Invalid nil pointer")
	}
	q.bridges = append(q.bridges, bridge)
	return nil
}

func (q *messageQueueStruct) ListMessageBridges() []*MessageBridge {
	return q.bridges
}

func (q *messageQueueStruct) Start() error {
	q.started = true
	go func() {
		defer func() {
			recover()
		}()
		//Read from internal receiver channel and transmit/dipatch/broadcast message to out/inout bridges via channel
		for q.started {
			select {
			case msg := <-q.internalBridgeChannel:
				go func(q *messageQueueStruct, incomingMsg common.Message) {
					defer func() {
						recover()
					}()
					for _, bridge := range q.bridges {
						if bridge != nil {
							if (*bridge).GetBridgeType() == WriteOnly || (*bridge).GetBridgeType() == ReadWrite {
								*((*bridge).GetInOutChannel()) <- incomingMsg
							}
						}
					}
				}(q, msg)
			case <-time.After(q.internalTimeOut):
				if logger != nil {
					logger.Debug(fmt.Sprintf("MessageQueue::TransmitOutbound::warn Internal timeout of %8.4f s reached!!", (float64(q.internalTimeOut)/float64(time.Second)))))
				}
			}
		}
	}()
	go func() {
		defer func() {
			recover()
		}()
		for q.started {
			//Read messages from bridges and send it to the internal receiver channel for transmit/dipatch/broadcast operations
			for _, bridge := range q.bridges {
				if bridge != nil {
					if (*bridge).GetBridgeType() == ReadOnly || (*bridge).GetBridgeType() == ReadWrite {
						go func(q *messageQueueStruct, bridge *MessageBridge) {
							select {
							case msg := <-*((*bridge).GetInOutChannel()):
								q.internalBridgeChannel <- msg
							case <-time.After(q.TimeOut):
								if logger != nil {
									logger.Debug(fmt.Sprintf("MessageQueue::ReceiveInbound::warn Timeout for bridge: %s of  %8.4f s reached!!", (*bridge).Name(), (float64(q.TimeOut) / float64(time.Second))))
								}
							}
						}(q, bridge)
					}
				}
			}
		}
	}()
	return nil
}

func (q *messageQueueStruct) Stop() error {
	return nil
}

func (q *messageQueueStruct) IsRunning() bool {
	return q.started
}

//Creates a New Message Queue
//
// Parameters:
//   messageReadTimeout (time.Duration) Timeout before reset the listening on each of Bridges (infinite if <= 0)
// Returns:
//   Message Queue New brand element
func NewMessageQueue(messageReadTimeout time.Duration) MessageQueue {
	var q messageQueueStruct = messageQueueStruct{
		TimeOut: messageReadTimeout,
	}
	runtime.SetFinalizer(&q, func(qs *messageQueueStruct) {
		defer func() {
			err := recover()
			if logger != nil {
				logger.Error(err.(error))
			}
		}()
		qs.Stop()
		qs.bridges = qs.bridges[:0]
		close(qs.internalBridgeChannel)
	})
	q.init()
	return &q

}
