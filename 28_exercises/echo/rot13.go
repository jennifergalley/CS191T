package main 

import (
	"bufio"
	"log"
	"net"
	"fmt"
	"io"
	"strings"
)

/*
Modify ‘echo’ so that it returns messages as rot 13
*/
func handle (conn net.Conn) {
	defer conn.Close()
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		
		if len(fs) < 1 {
			continue
		}
		
		switch fs[0] {
			case "echo":
				for i := 1; i < len(fs); i++ {
					word := fs[i]
					for j := 0; j < len(word); j++ {
						x := int32(word[j])
						x += 13
						fmt.Fprintf (conn, "%s", string(x))
					}
				}
				io.WriteString(conn, "\n")
			default:
				continue
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