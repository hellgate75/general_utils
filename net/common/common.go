package common

import (
	"github.com/hellgate75/general_utils/log"
	"net"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("net/common")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Type that describes Server Type
type ServerType int

const (
	// TCP Server Type
	TCP ServerType = iota + 1
	// UDP Protocol Server Type
	UDP
	//	// REST Server Type
	REST ServerType = iota + 1
	// Web Content Server Type
	CONTENT ServerType = iota + 1
)

// Type that describes Server State
type ServerStateSignal int

const (
	// RUNNING Server State - Server Operates Nornally
	RUNNING ServerStateSignal = iota + 500
	// PAUSED Server State - Server paused computation
	PAUSED
	// DISCONNECTING Server State - Server is disconnecting from client
	DISCONNECTING
	// REFUSED Server State - Server is refusing client connection
	REFUSED
	// FAILURE Server State - Server suspended computation due to Internal Server Error
	FAILURE
	// STOPPING Server State - Server stopping computation
	STOPPING
	// STOPPED Server State - Server stopped computation, ready for shutdown
	STOPPED
)

// Structure that describes Server Context
type ServerContext struct {
	Client ClientRef
	Server Server
}

// Type that describes Connection Handling Function
type ServerHablerFunc func(net.Conn, ServerContext, ...interface{}) error

// Type that describes Port Number
type Port int

// Structure that describes Port Numbers Interval
type PortInterval struct {
	// Interval Start Port Included
	Start Port
	// Interval End Port Included
	End Port
}

// Structure that describes Server Configuration
type ServerConfig struct {
	// Listening Port
	ListeningPort Port
	// Interval of Follow Up Ports for connected clients
	FollowUpPorts PortInterval
	// Log Verbosity Level
	LogLevel LogLevel
}

// Structure that describes Client Information
type ClientRef struct {
	// Client Address Informations
	Address IPAddressInfo
	// Client Hostname if provided
	Hostname string
	// Network Protocol
	Protocol string
	// Client Connection Port
	ClientPort Port
	// Server Follow Up Port
	ServerPort Port
	// Flags that define when a Client still seems connected
	IsAlive bool
	// Client connection
	Connection net.Conn
}
