package main

import (
	"bufio"
	"fmt"
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
	fmt.Printf("read bytes '%s'\n", string(bytes))
	if len(bytes) < 3 && bytes[0] == '\r' {
		// list contents
		c.Write([]byte("0The gopher RFC\trfc1436.txt\t127.0.0.1\t70"))
		c.Write([]byte("\r\n.\r\n"))
		c.Close()
	} else if string(bytes) == "rfc1436.txt\r\n" {
		c.Write([]byte("foobar\r\n.\r\n"))
		c.Close()
	}
}