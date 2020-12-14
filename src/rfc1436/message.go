package rfc1436

import (
	"fmt"
	"net"
)

type Listing struct  {
	Type Datatype
	Name string
	Location string
	Addr Address
}

type Address struct {
	Hostname string
	Port int
}

func (l Listing) Bytes() []byte {
	return []byte(fmt.Sprintf("%s%s\t%s\t%s\t%d\n", l.Type, l.Name, l.Location, l.Addr.Hostname, l.Addr.Port))
}

func NewError(err error) Listing {
	return Listing{
		Type: Error,
		Name: err.Error(),
	}
}

func WriteError(err error, c net.Conn) {
	c.Write(NewError(err).Bytes())
	c.Write([]byte(EndSentinel))
	c.Close()
}