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

// Interface describes MessageBroker features
type NoTransitionMessageBroker interface {
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

type noTransMessageBroker struct {
	sync.RWMutex
	inOutChan   chan common.Type
	initialized bool
	TimeOut     time.Duration
}

func (q *noTransMessageBroker) init() {
	if q.initialized {
		return
	}
	var IOchan chan common.Type = make(chan common.Type)
	q.inOutChan = IOchan
	q.initialized = true
}

func (q *noTransMessageBroker) Write(m common.Type) error {
	q.Lock()
	q.inOutChan <- m
	q.Unlock()
	return nil
}

func (q *noTransMessageBroker) Read() (common.Type, error) {
	var val common.Type
	if q.TimeOut == 0 {
		select {
		case msg := <-q.inOutChan:
			val = msg
		}
	} else {
		select {
		case msg := <-q.inOutChan:
			val = msg
		case <-time.After(q.TimeOut):
			return nil, errors.New("MessageBroker::Read::err Timeout reached")
		}
	}
	return val, nil
}

//Creates a New MessageBroker
//
// Parameters:
//   messageReadTimeout (time.Duration) Timeout before reset the listening (infinite if <= 0)
// Returns:
//   MessageBroker New brand queue
func NewNoTransitionMessageBroker(messageReadTimeout time.Duration) NoTransitionMessageBroker {
	var q noTransMessageBroker = noTransMessageBroker{
		TimeOut: messageReadTimeout,
	}
	runtime.SetFinalizer(&q, func(qs *noTransMessageBroker) {
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

// Interface describes Message MessageBroker features
type MessageBroker interface {
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

type messageBrokerStruct struct {
	sync.RWMutex
	internalBridgeChannel chan common.Message
	bridges               []*MessageBridge
	initialized           bool
	started               bool
	TimeOut               time.Duration
	internalTimeOut       time.Duration
}

func (q *messageBrokerStruct) init() {
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

func (q *messageBrokerStruct) AddMessageBridge(bridge *MessageBridge) error {
	if bridge == nil {
		return errors.New("MessageBroker::AddMessageBridge::err Invalid nil pointer")
	}
	q.bridges = append(q.bridges, bridge)
	return nil
}

func (q *messageBrokerStruct) ListMessageBridges() []*MessageBridge {
	return q.bridges
}

func (q *messageBrokerStruct) Start() error {
	q.started = true
	go func() {
		defer func() {
			recover()
		}()
		//Read from internal receiver channel and transmit/dipatch/broadcast message to out/inout bridges via channel
		for q.started {
			select {
			case msg := <-q.internalBridgeChannel:
				go func(q *messageBrokerStruct, incomingMsg common.Message) {
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
					logger.Debug(fmt.Sprintf("MessageBroker::TransmitOutbound::warn Internal timeout of %8.4f s reached!!", (float64(q.internalTimeOut) / float64(time.Second))))
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
						go func(q *messageBrokerStruct, bridge *MessageBridge) {
							select {
							case msg := <-*((*bridge).GetInOutChannel()):
								q.internalBridgeChannel <- msg
							case <-time.After(q.TimeOut):
								if logger != nil {
									logger.Debug(fmt.Sprintf("MessageBroker::ReceiveInbound::warn Timeout for bridge: %s of  %8.4f s reached!!", (*bridge).Name(), (float64(q.TimeOut) / float64(time.Second))))
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

func (q *messageBrokerStruct) Stop() error {
	q.started = false
	time.Sleep(time.Second)
	return nil
}

func (q *messageBrokerStruct) IsRunning() bool {
	return q.started
}

//Creates a New Message MessageBroker
//
// Parameters:
//   messageReadTimeout (time.Duration) Timeout before reset the listening on each of Bridges (infinite if <= 0)
// Returns:
//   Message MessageBroker New brand element
func NewMessageBroker(messageReadTimeout time.Duration) MessageBroker {
	var q messageBrokerStruct = messageBrokerStruct{
		TimeOut: messageReadTimeout,
	}
	runtime.SetFinalizer(&q, func(qs *messageBrokerStruct) {
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

type echoBridgeStruct struct {
	inOutBridgeChan chan common.Message
	echoBridgeChan  *chan common.Message
	running         bool
	initialized     bool
	timeout         time.Duration
	echoTimeout     time.Duration
}

func (b *echoBridgeStruct) init(echoBridgeChan *chan common.Message, timeout time.Duration, echoTimeout time.Duration) {
	if b.initialized {
		return
	}
	b.inOutBridgeChan = make(chan common.Message)
	b.echoBridgeChan = echoBridgeChan
	b.initialized = true
	b.running = false
	b.timeout = timeout
	b.echoTimeout = echoTimeout
}

func (b *echoBridgeStruct) Name() string {
	return "Echo R/W Bridge"
}
func (b *echoBridgeStruct) Open() error {
	b.running = true
	//Reading
	go func() {
		for b.running {
			select {
			case msg := <-b.inOutBridgeChan:
				*b.echoBridgeChan <- msg
			case <-time.After(b.timeout):
				if logger != nil {
					logger.Debug(fmt.Sprintf("EchoTestBridge::ReadFromChannel::warn Timeout <%d s> reached!!", b.timeout))
				}
			}
		}
	}()
	//Writing
	go func() {
		for b.running {
			select {
			case msg := <-*b.echoBridgeChan:
				b.inOutBridgeChan <- msg
			case <-time.After(b.echoTimeout):
				if logger != nil {
					logger.Debug(fmt.Sprintf("EchoTestBridge::ReadFromEchoChannel::warn Timeout <%d s> reached!!", b.echoTimeout))
				}
			}
		}
	}()
	return nil
}
func (b *echoBridgeStruct) Close() error {
	b.running = false
	return nil
}
func (b *echoBridgeStruct) IsOpen() bool {
	return b.running
}
func (b *echoBridgeStruct) GetInOutChannel() *chan common.Message {
	return &b.inOutBridgeChan
}
func (b *echoBridgeStruct) GetBridgeType() BridgeType {
	return ReadWrite
}

//Creates a New Message MessageBroker
//
// Parameters:
//   messageReadTimeout (time.Duration) Timeout before reset the listening on each of Bridges (infinite if <= 0)
// Returns:
//   Message MessageBroker New brand element
func NewEchoMessageBridge(echoBridgeChan *chan common.Message, messageReadTimeout time.Duration, echoMessageReadTimeout time.Duration) MessageBridge {
	var bridge echoBridgeStruct = echoBridgeStruct{}
	runtime.SetFinalizer(&bridge, func(qs *echoBridgeStruct) {
		defer func() {
			err := recover()
			if logger != nil {
				logger.Error(err.(error))
			}
		}()
		qs.Close()
		close(qs.inOutBridgeChan)
	})
	bridge.init(echoBridgeChan, messageReadTimeout, echoMessageReadTimeout)
	return &bridge

}
