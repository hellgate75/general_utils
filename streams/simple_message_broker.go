package streams

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/errors"
	"time"
)

type pipelineStruct struct {
	_out         chan common.Type
	_initialized bool
}

// Interface describes Simple Message Broker features
type SimpleMessageBroker interface {
	// Gets output channel.
	//
	// Returns:
	//   *chan common.Type Output channel pointer
	GetOutChannel() *chan common.Type
	// Writes item into the channel.
	//
	// Parameters:
	//   item (common.Type) Element to send  to the pipeline
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Write(common.Type) error
	// Reads item from the channel.
	//
	// Parameters:
	//   timeout (time.Duration) timeout before stop listening (0 or less means no timeout)
	//
	// Returns:
	//   common.Type Element read from the channel
	//   error Any suitable error risen during code execution
	Read(time.Duration) (common.Type, error)
	// Opens pipeline channel.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Open() error
	// Closes pipeline channel.
	//
	// Returns:
	//   error Any suitable error risen during code execution
	Close() error
	// Get pipeline status.
	//
	// Returns:
	//   bool Simple Message Broker opening state
	IsOpen() bool
}

func (this *pipelineStruct) GetOutChannel() *chan common.Type {
	return &this._out
}

func (this *pipelineStruct) Read(timeout time.Duration) (common.Type, error) {
	if !this._initialized {
		return nil, errors.New("Simple Message Broker not opened")
	}
	if timeout > 0 {
		select {
		case res := <-this._out:
			return res, nil
		case <-time.After(timeout):
			return nil, errors.New(fmt.Sprintf("Timeout %v reached, stop reading messages", timeout))
		}
	} else {
		select {
		case res := <-this._out:
			return res, nil
		}
	}
}

func (this *pipelineStruct) Write(item common.Type) error {
	if !this._initialized {
		return errors.New("Simple Message Broker not opened")
	}
	this._out <- item
	return nil
}

func (this *pipelineStruct) Open() error {
	if this._initialized {
		return errors.New("Simple Message Broker already opened")
	}
	this._out = make(chan common.Type)
	this._initialized = true
	return nil
}

func (this *pipelineStruct) Close() error {
	if !this._initialized {
		return errors.New("Simple Message Broker not opened")
	}
	close(this._out)
	this._initialized = false
	return nil
}

func (this *pipelineStruct) IsOpen() bool {
	return this._initialized
}

// Creates new pipeline.
//
// Returns:
//   streams.SimpleMessageBroker Any suitable error risen during code execution
func NewSimpleMessageBroker() SimpleMessageBroker {
	return &pipelineStruct{
		_initialized: false,
	}
}
