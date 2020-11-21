package main

import (
	"bufio"
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"github.com/raidancampbell/sonya/rfc1436"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
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
	if len(bytes) == 2 && bytes[0] == '\r' && bytes[1] == '\n' {
		// list contents
		listings, err := genListing("./serve")
		if err != nil {
			panic(err)
		}
		for _, listing := range listings {
			c.Write([]byte(listing.String()))
		}
		c.Write([]byte("\r\n.\r\n"))
		c.Close()
	} else {
		request := strings.TrimSuffix(string(bytes), "\r\n")
		sub, err := isSubDir("./serve", request)
		if err != nil {
			panic(err)
		}
		if !sub {
			fmt.Printf("attempted to read from outside serve directory! '%s'\n", request)
			c.Close()
			return
		}
		stat, err := os.Stat(request)
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			// TODO: this is duplicated
			listings, err := genListing(request)
			if err != nil {
				panic(err)
			}
			for _, listing := range listings {
				c.Write([]byte(listing.String()))
			}
			c.Write([]byte("\r\n.\r\n"))
			c.Close()
		} else {
			s, err := rstrings.FileToString(request)
			if err != nil {
				panic(err)
			}
			c.Write([]byte(s))
			c.Write([]byte("\r\n.\r\n"))
			c.Close()
		}
	}
}

func genListing(basePath string) ([]rfc1436.Listing, error) {
	var listings []rfc1436.Listing

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return listings, err
	}

	for _, file := range files {
		if file.IsDir() {
			listings = append(listings, rfc1436.Listing{
				Type:     rfc1436.Directory,
				Name:     file.Name(),
				Location: filepath.Join(basePath, file.Name()),
				Addr:     rfc1436.Address{
					Hostname: "127.0.0.1",
					Port: 7000,
				},
			})
		} else {
			listings = append(listings, rfc1436.Listing{
				Type:     rfc1436.Binary, // TODO: everything is a binary file. correctly recognize types
				Name:     file.Name(),
				Location: filepath.Join(basePath, file.Name()),
				Addr:     rfc1436.Address{
					Hostname: "127.0.0.1",
					Port: 7000,
				},
			})
		}
	}
	return listings, nil
}

// thanks, https://stackoverflow.com/a/62529061
func isSubDir(parent, sub string) (bool, error) {
	up := ".." + string(os.PathSeparator)

	// path-comparisons using filepath.Abs don't work reliably according to docs (no unique representation).
	rel, err := filepath.Rel(parent, sub)
	if err != nil {
		return false, err
	}
	if !strings.HasPrefix(rel, up) && rel != ".." {
		return true, nil
	}
	return false, nil
}
