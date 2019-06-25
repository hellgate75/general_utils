package net

import (
	"fmt"
	"net"
)

type HandlerFunc func(net.Conn) error

type ServerType int

const (
	TCP ServerType = iota + 1
	IP ServerType
	UDP ServerType
)

type Server interface {
	Open() error
	Close() error
}

func InitServers() {
	fmt.Println(fmt.Sprintf("%v", TCP))
	fmt.Println(fmt.Sprintf("%v", IP))
	fmt.Println(fmt.Sprintf("%v", UDP))
}
