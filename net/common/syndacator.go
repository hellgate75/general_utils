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
	//    ServerHablerFunc Represents one of client connection consumer to be removed
	// Returns:
	//    error Any error that can occurs during computation
	Deallocate(client ClientRef) error
	// Remove Connection handling behaviour from list
	// Returns:
	//    error Any error that can occurs during computation
	List() []ClientRef
	// Remove Connection handling behaviour from list
	// Returns:
	//    error Any error that can occurs during computation
	CheckAlives() error
}

type __portSyndacatorStruct struct {
	__clientMap    map[Port]ClientRef
	__portInterval PortInterval
}

func (syn *__portSyndacatorStruct) Allocate(conn net.Conn) (*ClientRef, error) {
	addr := conn.LocalAddr()
	addrInfo, err := GetIpAddressInfo(addr.String())
	if err != nil {
		return nil, err
	}
	var host string = ""
	if names, err := net.LookupAddr(addrInfo.Ip.String()); err == nil && len(names) > 0 {
		host = names[0]
	}
	port := syn.__portInterval.Start
	for _, ok := syn.__clientMap[port]; ok; {
		port = Port(int(port) + 1)
	}
	if port < syn.__portInterval.Start || port > syn.__portInterval.End {
		return nil, errors.New(fmt.Sprintf("None available port: Server Busy!! Provisioned port <%s> out of Range <%v, %v>", port, syn.__portInterval.Start, syn.__portInterval.End))
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
	syn.__clientMap[port] = cr
	return &cr, nil
}

func (syn *__portSyndacatorStruct) Deallocate(client *ClientRef) error {
	if client == nil {
		return errors.New("Client Reference cannot be Nil")
	}
	delete(syn.__clientMap, client.ServerPort)
	return nil
}

func (syn *__portSyndacatorStruct) List() []ClientRef {
	var outList []ClientRef = []ClientRef{}
	for _, val := range syn.__clientMap {
		outList = append(outList, val)
	}
	return outList
}
