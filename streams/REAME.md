package streams // import "github.com/hellgate75/general_utils/streams"


FUNCTIONS

func DownloadFile(filepath string, url string) error
    Dowload file from url and save locally.

    Parameters:

    filepath (string) Absolute destination file path
    url (string) Source file url

    Returns:

    error Any suitable error risen during code execution

func DownloadFileAsByteArray(url string) ([]byte, error)
    Dowload file from url and return content byte array.

    Parameters:

    url (string) Source file url

    Returns:

    []byte Bytes contained in the remote support
    error Any suitable error risen during code execution

func GetCurrentPath() string
    Get current Go execution path.

    Returns:

    string Current absolute path

func InitLogger()
func LoadFileBytes(path string) ([]byte, error)
    Load all content in a file as byte array.

    Parameters:

    path (string) Absolute file path

    Returns:

    []byte File content
    error Any suitable error risen during code execution

func LoadFileContent(path string) (string, error)
    Load all content in a file as string.

    Parameters:

    path (string) Absolute file path

    Returns:

    string File content
    error Any suitable error risen during code execution


TYPES

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

func NewEchoMessageBridge(echoBridgeChan *chan common.Message, messageReadTimeout time.Duration, echoMessageReadTimeout time.Duration) MessageBridge
    Creates a New Message MessageBroker

    Parameters:

    messageReadTimeout (time.Duration) Timeout before reset the listening on each of Bridges (infinite if <= 0)

    Returns:

    Message MessageBroker New brand element

type MessageBroker interface {

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
	// Has unexported methods.
}
    Interface describes Message MessageBroker features

func NewMessageBroker(messageReadTimeout time.Duration) MessageBroker
    Creates a New Message MessageBroker

    Parameters:

    messageReadTimeout (time.Duration) Timeout before reset the listening on each of Bridges (infinite if <= 0)

    Returns:

    Message MessageBroker New brand element

type NoTransitionMessageBroker interface {

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
	// Has unexported methods.
}
    Interface describes MessageBroker features

func NewNoTransitionMessageBroker(messageReadTimeout time.Duration) NoTransitionMessageBroker
    Creates a New MessageBroker

    Parameters:

    messageReadTimeout (time.Duration) Timeout before reset the listening (infinite if <= 0)

    Returns:

    MessageBroker New brand queue

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
    Interface describes Simple Message Broker features

func NewSimpleMessageBroker() SimpleMessageBroker
    Creates new pipeline.

    Returns:

    streams.SimpleMessageBroker Any suitable error risen during code execution

