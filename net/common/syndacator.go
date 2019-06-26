package common

import (
	"errors"
	"fmt"
	"net"
)

type PortSyndacator interface {
	// Allocate New Client
	// Parameters:
	//    net.Conn Client Connection
	// Returns:
	//    (*ClientRef Pointer to just added client connection,
	//    error Any error that can occurs during computation)
	Allocate(net.Conn) (*ClientRef, error)
	// Deallocate Current Client
	// Parameters:
	//    client (*ClientRef) Represents pointer to Element to Deallocate
	// Returns:
	//    error Any error that can occurs during computation
	Deallocate(client *ClientRef) error
	// List of Stored Clients
	// Returns:
	//    []ClientRef List of Stored Client Reference Information
	List() []ClientRef
	// Remove Connection handling behaviour from list
	// Parameetrs:
	//   port (Port) Required Client Port
	// Returns:
	//    (*ClientRef Required Client Reference Pointer,
	//    error Any error that can occurs during computation)
	ClientByPort(port Port) (*ClientRef, error)
}

type __portSyndacatorStruct struct {
	ClientsMap   map[Port]ClientRef
	PortInterval PortInterval
}

func (syn *__portSyndacatorStruct) Allocate(conn net.Conn) (*ClientRef, error) {
	addr := conn.RemoteAddr()
	addrInfo, err := GetIpAddressInfo(addr.String())
	if err != nil {
		return nil, err
	}
	var host string = ""
	if names, err := net.LookupAddr(addrInfo.Ip.String()); err == nil && len(names) > 0 {
		host = names[0]
	}
	port := syn.PortInterval.Start
	for _, ok := syn.ClientsMap[port]; ok; {
		port = Port(int(port) + 1)
	}
	if port < syn.PortInterval.Start || port > syn.PortInterval.End {
		return nil, errors.New(fmt.Sprintf("None available port: Server Busy!! Provisioned port <%v> out of Range <%v, %v>", port, syn.PortInterval.Start, syn.PortInterval.End))
	}
	cr := ClientRef{
		Address:    *addrInfo,
		Hostname:   host,
		Protocol:   addr.Network(),
		ClientPort: 0,
		ServerPort: port,
		IsAlive:    true,
		Connection: conn,
	}
	syn.ClientsMap[port] = cr
	return &cr, nil
}

func (syn *__portSyndacatorStruct) ClientByPort(port Port) (*ClientRef, error) {
	if cli, ok := syn.ClientsMap[port]; ok {
		return &cli, nil
	} else {
		return nil, errors.New(fmt.Sprintf("No Client Allocated at Port <%v>!!", port))
	}

}

func (syn *__portSyndacatorStruct) Deallocate(client *ClientRef) error {
	if client == nil {
		return errors.New("Client Reference cannot be Nil")
	}
	delete(syn.ClientsMap, client.ServerPort)
	return nil
}

func (syn *__portSyndacatorStruct) List() []ClientRef {
	var outList []ClientRef = []ClientRef{}
	for _, val := range syn.ClientsMap {
		outList = append(outList, val)
	}
	return outList
}

func NewPortSyndacator(interval PortInterval) PortSyndacator {
	return &__portSyndacatorStruct{
		make(map[Port]ClientRef),
		interval,
	}
}
