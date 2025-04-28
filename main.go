package main

import (
	"fmt"
	"net"
)

func main() {
	ParseHTTPRequest("guh?")

	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:2100")
	if err != nil { // unrecoverable
		panic(err)
	}

	ln, err := net.ListenTCP("tcp", laddr) // TODO: change to have arbitrary port from commandline args
	if err != nil {                        // unrecoverable
		panic(err)
	}
	defer ln.Close()

	for {
		handleClient(ln)
	}
}

func handleClient(ln *net.TCPListener) {
	conn, err := ln.Accept()

	if err != nil {
		fmt.Println("guh!")
		panic(err)
	}
	defer conn.Close()

	b, err := readClientData(conn)
	if err != nil {
		fmt.Println("guh!!")
		panic(err)
	}

	ParseHTTPRequest(string(b))
}

func readClientData(conn net.Conn) ([]byte, error) {
	// TODO: change to only read up until HTTP terminator
	b := make([]byte, 2048)
	nbytes, err := conn.Read(b)
	if err != nil {
		return b, err
	}
	return b[:nbytes], nil
}
