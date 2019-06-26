package common

import ()

// Server Interface that describes Server features
type Server interface {
	// Open Server Listener
	// Returns:
	//    error Any error that can occurs during computation
	Open() error
	// Close Server Listener and consume pending client connections
	// Returns:
	//    error Any error that can occurs during computation
	Close() error
	// Check if Server is Listeing for Clients Requests
	// Returns:
	//    bool Running state
	IsListening() bool
	// Check if Server is still Opened serving clients
	// Returns:
	//    bool Open state
	IsRunning() bool
	// Streams an interface to the connected clients
	// Parameters:
	//    interface{} Represents the element must be sent at the clients in Broadcast
	// Returns:
	//    error Any error that can occurs during computation
	Stream(interface{}) error
	// Reads first element coming from any of connected clients
	// Returns:
	//    (interface{} Represents the element read from one of connected clients,
	//    error Any error that can occurs during computation)
	Receive() (interface{}, error)
	// Add Connection handling behaviour behalf the given function
	// Parameters:
	//    ServerHablerFunc Represents one of client connection consumer to be added
	// Returns:
	//    error Any error that can occurs during computation
	HandleConnOn(ServerHablerFunc) error
	// List Enabled Connection handling functions
	// Returns:
	//    []ServerHablerFunc List of active handlers
	HandlingFuncs() []ServerHablerFunc
	// Remove Connection handling behaviour from list
	// Parameters:
	//    ServerHablerFunc Represents one of client connection consumer to be removed
	// Returns:
	//    error Any error that can occurs during computation
	RevokeHandler(ServerHablerFunc) error
	// Clear Server internal references and eventually Close the Server in case it's Open/Listening
	// Returns:
	//    error Any error that can occurs during computation
	Destroy() error
	// Returns list of connected Clients
	Clients() []ClientRef
	// Log on specific interface channel
	// Parameters:
	//   channel (chan interface{}) Output channel for server logging activities
	LogOn(channel chan interface{})
}
