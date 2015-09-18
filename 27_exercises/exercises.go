package main

import (
	"net"
	"log"
	"bufio"
	"io"
)

/*
Create an ‘echo’ TCP server. 
It should accept a connection and write to the connection anything that’s sent to it.
*/
func handle (conn net.Conn) {
	defer conn.Close()
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text() + "\n"
		io.WriteString(conn, ln)
	}
}

// telnet localhost 9000 to connect to it.

func main () {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
	
}