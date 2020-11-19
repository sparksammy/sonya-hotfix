package main

import (
	"bufio"
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"github.com/raidancampbell/sonya/rfc1436"
	"net"
	"time"
)

func main() {

	ln, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
func handleConnection(c net.Conn) {

	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	br := bufio.NewReader(c)

	bytes, err := br.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("read bytes '%v'\n", bytes)
	if len(bytes) < 3 && bytes[0] == '\r' {
		// list contents
		listing := rfc1436.Listing{
			Type:     rfc1436.Text,
			Location: "rfc1436.txt",
			Addr:     rfc1436.Address{
				Hostname: "127.0.0.1",
				Port: 7000,
			},
		}
		c.Write([]byte(listing.String()))
		c.Write([]byte("\r\n.\r\n"))
		c.Close()
	} else if string(bytes) == "rfc1436.txt\r\n" {
		s, err := rstrings.FileToString("rfc1436.txt")
		if err != nil {
			panic(err)
		}
		c.Write([]byte(s))
		c.Write([]byte("\r\n.\r\n"))
		c.Close()
	}
}