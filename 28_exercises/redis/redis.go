package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"io"
	"fmt"
)

/*
Create a simplified redis clone which accepts GET, SET and DEL commands. 
‘GET <KEY>’ should write the value of <KEY> followed by a new line. 
‘SET <KEY> <VALUE>’ should set the value of <KEY>. 
‘DEL <KEY>’ should remove the value. 
Data should be stored in memory.
*/
func handle (conn net.Conn) {
	defer conn.Close()
	
	data := map[string]string {"":""}
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		
		if len(fs) < 1 {
			continue
		}
		
		switch fs[0] {
			case "GET":
				key := fs[1]
				value := data[key]
				fmt.Println("Get:", key, value)
				fmt.Fprintf (conn, "%s\n", value)
			case "SET":
				if len(fs) != 3 {
					io.WriteString(conn, "Invalid number of arguments. Expected: 2\n")
					continue
				}
				key := fs[1]
				value := fs[2]
				data[key] = value
				fmt.Println("Set:", key, value)
			case "DEL":
				key := fs[1]
				delete(data, key)
				fmt.Println("Deleted:", key)
			default:
				io.WriteString(conn, "Invalid command\n")
		}
	}
}


func main () {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
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