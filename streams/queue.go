package streams

import (
	"github.com/hellgate75/general_utils/common"
)

// Interface describes Queue features
type Queue interface {
	// Gets output channel.
	//
	// Returns:
	//   *chan common.Type Output channel pointer
	GetOutChannel() *chan common.Type
	//Starts the pipeline flow
	Start()
	//Stops the pipeline flow
	Stop()
	// Gets the pipeline running state.
	//
	// Returns:
	//   bool Running state
	IsRunning() bool
	// Writes item into the channel.
	//
	// Parameters:
	//   url (string) Source file url
	Write(common.Type)
	// Reads item from the channel.
	//
	// Returns:
	//   common.Type Element read from the channel
	//   error Any suitable error risen during code execution
	Read() (common.Type, error)
}
