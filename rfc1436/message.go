package rfc1436

import "fmt"

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

func (l Listing) String() string {
	return fmt.Sprintf("%s%s%s\t%s\t%d", l.Type, l.Name, l.Location, l.Addr.Hostname, l.Addr.Port)
}